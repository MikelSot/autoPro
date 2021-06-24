package dto

import "time"

// para mostra al cliente
type InvoiceClients []*InvoiceClients
type InvoiceClient struct {
	Ruc             string    `json:"ruc"`
	Status          string    `json:"status"`
	InvoiceDate     time.Time `json:"invoice_date"`
}

// para mostra al mecanico
type InvoiceWorkshops []*InvoiceWorkshop
type InvoiceWorkshop struct {
	Ruc             string    `json:"ruc"`
	Status          string    `json:"status"`
	InvoiceDate     time.Time `json:"invoice_date"`
	EmployeeID      uint      `json:"employee_id"`
	PaymentMethodID uint8     `json:"payment_method_id"`
}

