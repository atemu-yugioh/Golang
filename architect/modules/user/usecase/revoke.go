package usecase

import (
	"context"

	"github.com/google/uuid"
)

// business logic

type revokeUC struct {
	sessionCommandRepo SessionCommandRepository
}

func NewRevokeUC(sessionCommandRepo SessionCommandRepository) *revokeUC {
	return &revokeUC{
		sessionCommandRepo: sessionCommandRepo,
	}
}

func (uc *revokeUC) RevokeToken(ctx context.Context, id uuid.UUID) error {
	if err := uc.sessionCommandRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
