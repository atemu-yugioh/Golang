package biz

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

type DeleteTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
	DeleteTodo(ctx context.Context, cond map[string]interface{}) error
}

type deleteTodoBiz struct {
	store DeleteTodoStorage
}

func NewDeleteTodoBiz(store DeleteTodoStorage) *deleteTodoBiz {
	return &deleteTodoBiz{store}
}

func (biz *deleteTodoBiz) DeleteTodoById(ctx context.Context, id int) error {
	todoFound , err :=  biz.store.GetTodo(ctx,map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName,err)
		}

		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if todoFound.Status != nil && *todoFound.Status == model.ItemStatusDeleted {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.DeleteTodo(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}