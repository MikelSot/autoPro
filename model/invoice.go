package model

import "time"

type Invoice struct {
	ID              uint      `gorm:"primaryKey" json:"id"` // ,omitempty
	Ruc             string    `gorm:"type:varchar(20); default:''; not null" json:"ruc"`
	Status          string    `gorm:"type:varchar(25); default:'pendiente'; not null" json:"status"`
	InvoiceDate     time.Time `gorm:"not null" json:"invoice_date"`
	ClientID        uint      `gorm:"not null" json:"client_id"`
	EmployeeID      uint      `gorm:"not null" json:"employee_id"`
	WorkshopID      uint8     `gorm:"not null" json:"workshop_id"`
	PaymentMethodID uint8     `gorm:"not null" json:"payment_method_id"`
	InvoiceItems    []InvoiceItem
	TimeModel
}

type Invoices []*Invoice