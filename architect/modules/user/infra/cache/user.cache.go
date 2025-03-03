package cache

import (
	"architect/modules/user/domain"
	"architect/modules/user/usecase"
	"context"
)

// Proxy design pattern
type userCacheRepo struct {
	realRepo usecase.UserQueryRepository
	cache    map[string]*domain.User
}

func NewUserCacheRepo(realRepo usecase.UserQueryRepository, cache map[string]*domain.User) userCacheRepo {
	return userCacheRepo{
		realRepo: realRepo,
		cache:    cache,
	}
}

func (c userCacheRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if user, ok := c.cache[email]; ok {
		return user, nil
	}

	user, err := c.realRepo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	c.cache[email] = user

	return user, nil
}
