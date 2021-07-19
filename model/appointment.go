package model

import "time"

type Appointment struct {
	ID               uint      `gorm:"primaryKey" json:"id"` // ,omitempty
	Workshop         string    `gorm:"type:varchar(60); default:''; not null" json:"workshop"`
	Service          string    `gorm:"type:varchar(100); default:''; not null" json:"service"`
	Description      string    `gorm:"type:varchar(350); default:''" json:"description"`
	DateHour         time.Time `gorm:"not null" json:"date_hour"`
	OrderAttention   string    `gorm:"type:char(2); default:'0''; not null" json:"order_attention"`
	State            string    `gorm:"type:varchar(30); default:'espera'" json:"state"`
	VehicleType      string    `gorm:"type:varchar(60); default:''; not null" json:"vehicle_type"`
	PickUp           bool      `gorm:"default:false; not null" json:"pick_up"`
	ClientID         uint      `gorm:"not null" json:"client_id"`
	TimeModel
}

type Appointments []*Appointment

