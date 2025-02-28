package domain

import "architect/common"

type Product struct {
	common.BaseModel
	CategoryId string `json:"category_id" gorm:"category_id;"`
	Name       string `json:"name" gorm:"name;"`
	// Image      any    `json:"image" gorm:"image;"`
	Type        string `json:"type" gorm:"type;"`
	Description string `json:"description" gorm:"description;"`
}
