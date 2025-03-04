package usecase

import (
	"architect/common"
	"architect/modules/user/domain"
	"context"
	"errors"
)

type registerUC struct {
	userQueryRepo UserQueryRepository
	userCmdRepo   UserCommandRepository
	hasher        Hasher
}

func NewRegisterUC(userQueryRepo UserQueryRepository, userCmdRepo UserCommandRepository, hasher Hasher) *registerUC {
	return &registerUC{userQueryRepo: userQueryRepo, userCmdRepo: userCmdRepo, hasher: hasher}
}

func (uc *registerUC) Register(ctx context.Context, dto EmailPasswordRegisterDTO) error {
	// 1. Find user by email:
	// 1.1 Found: return error (email has existed)
	// 2. Generate salt
	// 3. Hash password+salt
	// 4. Create user entity

	userFound, err := uc.userQueryRepo.FindByEmail(ctx, dto.Email)

	if userFound != nil {
		return domain.ErrEmailHasExisted
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return err
	}

	salt, err := uc.hasher.RandomStr(30)

	if err != nil {
		return err
	}

	hashedPassword, err := uc.hasher.HashPassword(salt, dto.Password)

	if err != nil {
		return err
	}

	userEntity, err := domain.NewUser(
		common.GenUUID(),
		dto.FirstName,
		dto.LastName,
		dto.Email,
		hashedPassword,
		salt,
		domain.RoleUser,
		"activated",
	)

	if err != nil {
		return err
	}

	if err := uc.userCmdRepo.Create(ctx, userEntity); err != nil {
		return err
	}

	return nil
}
