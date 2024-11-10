package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required,oneof=complete incomplete"`
}