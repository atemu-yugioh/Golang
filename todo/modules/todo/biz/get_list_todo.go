package biz

import (
	"context"
	"todos/common"
	"todos/modules/todo/model"
)

type ListTodoStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	)([]model.Todo, error)
}



type listTodoBiz struct {
	store ListTodoStorage
}

func NewListTodoBiz(store ListTodoStorage) *listTodoBiz {
	return &listTodoBiz{store}
}

func (biz *listTodoBiz) ListTodo(context context.Context, filter *model.Filter, paging *common.Paging) ([]model.Todo, error) {
	data, err := biz.store.ListItem(context, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}