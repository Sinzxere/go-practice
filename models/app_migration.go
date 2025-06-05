package models

import (
	"time"
)

type AppMigration struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AppName   string    `gorm:"size:255;not null" json:"app_name"`
	Migrated  bool      `gorm:"default:false" json:"migrated"`
	Restored  bool      `gorm:"default:false" json:"restored"`
	Recheck   bool      `gorm:"default:false" json:"recheck"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
