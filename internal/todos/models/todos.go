package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" validate:"required"`
	Status    string         `json:"status" validate:"required,oneof=complete incomplete"`
	Context   string         `json:"context" validate:"required"`                   
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`    
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`   
}
