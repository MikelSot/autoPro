package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrClientCanNotBeNill = errors.New("El cliente no puede ser nula")
	ErrIDClientDoesNotExists = errors.New("El cliente no existe")
)



type TimeModel struct {
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
