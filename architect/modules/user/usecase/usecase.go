package usecase

import (
	"architect/modules/user/domain"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegisterDTO) error
	Login(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
	RevokeToken(ctx context.Context, id uuid.UUID) error
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error)
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(passwordHold, salt, password string) bool
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpireInSeconds() int
	RefreshExpireInSeconds() int
}

type useCase struct {
	*registerUC
	*LoginUC
	*revokeUC
	*refreshTokenUC
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCommandRepo() UserCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCommandRepo() SessionCommandRepository
	BuildSessionRepo() SessionRepository
}

func UseCaseWithBuilder(b Builder) UseCase {

	return &useCase{
		registerUC:     NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCommandRepo(), b.BuildHasher()),
		LoginUC:        NewLoginUC(b.BuildUserQueryRepo(), b.BuildSessionCommandRepo(), b.BuildTokenProvider(), b.BuildHasher()),
		revokeUC:       NewRevokeUC(b.BuildSessionCommandRepo()),
		refreshTokenUC: NewRefreshTokenUC(b.BuildUserQueryRepo(), b.BuildSessionRepo(), b.BuildTokenProvider(), b.BuildHasher()),
	}

}

func NewUseCase(userRepo UserRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) UseCase {
	return &useCase{
		registerUC:     NewRegisterUC(userRepo, userRepo, hasher),
		LoginUC:        NewLoginUC(userRepo, sessionRepo, tokenProvider, hasher),
		revokeUC:       NewRevokeUC(sessionRepo),
		refreshTokenUC: NewRefreshTokenUC(userRepo, sessionRepo, tokenProvider, hasher),
	}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}

type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, user *domain.User) error
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCommandRepository
}

type SessionQueryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.Session, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.Session, error)
}

type SessionCommandRepository interface {
	Create(ctx context.Context, data *domain.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}
