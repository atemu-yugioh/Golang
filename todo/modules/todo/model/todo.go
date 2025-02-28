package model

import (
	"errors"
	"todos/common"
)

const (
	EntityName = "Todo"
)


var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrTodoDeleted = errors.New("todo is deleted")
)

type Todo struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (Todo) TableName() string { return "todos" }

type TodoCreate struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoCreate) TableName() string { return Todo{}.TableName() }

type TodoUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoUpdate) TableName() string { return Todo{}.TableName() }