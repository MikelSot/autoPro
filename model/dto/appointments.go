package dto

import "time"

type AppointmentCreate struct {
	Workshop         string    `json:"workshop"`
	Service          string    `json:"service"`
	Description      string    `json:"description"`
	DateHour         time.Time `json:"date_hour"`
	OrderAttention   string    `json:"order_attention"`
	VehicleType      string    `json:"vehicle_type"`
	PickUp           bool      `json:"pick_up"`
	ClientID         uint      `json:"client_id"`
}

type AppointmentUpdate struct {
	Workshop         string    `json:"workshop"`
	Service          string    `json:"service"`
	Description      string    `json:"description"`
	DateHour         time.Time `json:"date_hour"`
	OrderAttention   string    `json:"order_attention"`
	State            string    `json:"state"`
	VehicleType      string    `json:"vehicle_type"`
	PickUp           bool      `json:"pick_up"`
	ClientID         uint      `json:"client_id"`
}


type AppointmentUpdateState struct {
	State   string  `json:"state"`
}
