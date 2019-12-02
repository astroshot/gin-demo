package model

import (
	"time"
)

// BaseModel defines basic attributes
type BaseModel struct {
	ID        *int64     `gorm:"primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

// StatusModel extends BaseModel with attribute `Status`
type StatusModel struct {
	Status *int
	BaseModel
}
