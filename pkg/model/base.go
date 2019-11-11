package model

import (
	"time"
)

// BaseModel defines basic attributes
type BaseModel struct {
	ID        *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// StatusModel extends BaseModel with attribute `Status`
type StatusModel struct {
	Status *int
	BaseModel
}
