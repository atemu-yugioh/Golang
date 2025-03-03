package usecase

import (
	"architect/modules/user/domain"
	"context"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegisterDTO) error
	Login(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
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
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCommandRepo() UserCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCommandRepo() SessionCommandRepository
}

func UseCaseWithBuilder(b Builder) UseCase {

	return &useCase{
		registerUC: NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCommandRepo(), b.BuildHasher()),
		LoginUC:    NewLoginUC(b.BuildUserQueryRepo(), b.BuildSessionCommandRepo(), b.BuildTokenProvider(), b.BuildHasher()),
	}

}

func NewUseCase(userRepo UserRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) UseCase {
	return &useCase{
		registerUC: NewRegisterUC(userRepo, userRepo, hasher),
		LoginUC:    NewLoginUC(userRepo, sessionRepo, tokenProvider, hasher),
	}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}

type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, user *domain.User) error
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCommandRepository
}

type SessionQueryRepository interface{}

type SessionCommandRepository interface {
	Create(ctx context.Context, data *domain.Session) error
}
