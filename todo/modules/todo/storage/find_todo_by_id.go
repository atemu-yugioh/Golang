package storage

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error) {
	var data model.Todo

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}