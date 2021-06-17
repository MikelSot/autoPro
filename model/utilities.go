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

type FirstDataModel struct{
	ID      uint8  `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active  bool   `gorm:"default:true; not null" json:"active"`
}