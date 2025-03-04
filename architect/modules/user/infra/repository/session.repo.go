package repository

import (
	"architect/common"
	"architect/modules/user/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TbSessionName = "user_sessions"

type sessionMySQLRepo struct {
	db *gorm.DB
}

func NewSessionMySQLRepo(db *gorm.DB) sessionMySQLRepo {
	return sessionMySQLRepo{db: db}
}

func (repo sessionMySQLRepo) FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.Session, error) {

	var dto SessionDTO

	if err := repo.db.Table(TbSessionName).Where("refresh_token = ?", refreshToken).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}

func (repo sessionMySQLRepo) FindById(ctx context.Context, id uuid.UUID) (*domain.Session, error) {

	var dto SessionDTO

	if err := repo.db.Table(TbSessionName).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}

func (repo sessionMySQLRepo) Create(ctx context.Context, data *domain.Session) error {
	dto := SessionDTO{
		Id:           data.Id(),
		UserId:       data.UserId(),
		RefreshToken: data.RefreshToken(),
		AccessExpAt:  data.AccessExpAt(),
		RefreshExpAt: data.RefreshExpAt(),
	}

	return repo.db.Table(TbSessionName).Create(&dto).Error
}

func (repo sessionMySQLRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.Table(TbSessionName).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
