package model

import (
	"time"
)

// BaseModel defines basic attributes
type BaseModel struct {
	ID        *int64     `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updateAt"`
}

// StatusModel extends BaseModel with attribute `Status`
type StatusModel struct {
	Status *int `gorm:"status" json:"status"`
	BaseModel
}
