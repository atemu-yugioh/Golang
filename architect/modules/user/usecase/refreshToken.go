package usecase

import (
	"architect/common"
	"architect/modules/user/domain"
	"context"
	"errors"
	"fmt"
	"time"
)

type refreshTokenUC struct {
	userQueryRepo UserQueryRepository
	sessionRepo   SessionRepository
	tokenProvider TokenProvider
	hasher        Hasher
}

func NewRefreshTokenUC(userQueryRepo UserQueryRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) *refreshTokenUC {
	return &refreshTokenUC{userQueryRepo, sessionRepo, tokenProvider, hasher}
}

func (uc *refreshTokenUC) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error) {
	fmt.Println("Session", refreshToken)
	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)

	if err != nil {
		return nil, err
	}

	if session.RefreshExpAt().UnixNano() < time.Now().UTC().UnixNano() {
		return nil, errors.New("refresh token expired")
	}

	user, err := uc.userQueryRepo.FindById(ctx, session.UserId())

	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, errors.New("user has been banned")
	}

	userId := user.Id()
	sessionId := common.GenUUID()

	// 3. Gen JWT
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())

	if err != nil {
		return nil, err
	}

	// 4. insert new session into DB
	newRefreshToken, _ := uc.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpireInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.RefreshExpireInSeconds()))

	newSession := domain.NewSession(sessionId, userId, newRefreshToken, tokenExpAt, refreshExpAt)

	if err := uc.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	// 5. delete session invalid
	go func() {
		_ = uc.sessionRepo.Delete(ctx, session.Id())
	}()

	// 5. Return token response dto

	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:      newRefreshToken,
		RefreshTokenExpIn: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil
}
