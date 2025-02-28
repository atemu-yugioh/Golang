package storage

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

func (s *sqlStore) CreateITem(ctx context.Context, data *model.TodoCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}