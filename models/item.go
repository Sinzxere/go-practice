package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	Price       float64        `json:"price" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
