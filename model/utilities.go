package model

import (
	"gorm.io/gorm"
	"time"
)

type TimeModel struct {
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
