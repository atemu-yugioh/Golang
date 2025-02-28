package biz

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

type UpdateTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
	UpdateTodo(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoUpdate) error
}

type updateTodoBiz struct {
	store UpdateTodoStorage
}

func NewUpdateTodoBiz(store UpdateTodoStorage) *updateTodoBiz {
	return &updateTodoBiz{store}
}

func (biz *updateTodoBiz) UpdateTodoById(ctx context.Context, id int, dataUpdate *model.TodoUpdate) error {
	todoFound, err := biz.store.GetTodo(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName,err)
		}

		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if todoFound.Status != nil && *todoFound.Status == model.ItemStatusDeleted {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.UpdateTodo(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}