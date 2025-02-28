package storage

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

func (s *sqlStore) UpdateTodo(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}