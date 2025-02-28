package biz

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

type GetTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
}

type getTodoBiz struct {
	store GetTodoStorage
}

func NewGetTodoBiz(store GetTodoStorage) *getTodoBiz {
	return &getTodoBiz{store}
}

func (biz *getTodoBiz) GetTodoById(ctx context.Context, id int) (*model.Todo, error) {
	data, err := biz.store.GetTodo(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}

