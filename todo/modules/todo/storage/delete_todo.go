package storage

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)


func (s *sqlStore) DeleteTodo(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.Todo{}.TableName()).Where(cond).Updates(map[string]interface{}{"status": model.ItemStatusDeleted.String()}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}