package model

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"size:255" json:"title" validate:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
