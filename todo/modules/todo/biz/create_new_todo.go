package biz

import (
	"context"
	"strings"
	"todos/common"
	"todos/modules/todo/model"
)

type CreateItemStorage interface {
	CreateITem(ctx context.Context, data *model.TodoCreate) error
}

type createItemBiz struct {
	store	CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx	context.Context, data *model.TodoCreate) error {

	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.store.CreateITem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName,err)
	}

	return nil
}
