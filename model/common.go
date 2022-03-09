package model

import (
	"time"

	"gorm.io/gorm"
)

// CommonModel https://gorm.io/docs/models.html#gorm-Model
type CommonModel struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Page struct {
	PageNum  int `json:"page" form:"page"`
	PageSize int `json:"size" form:"size"`
}
