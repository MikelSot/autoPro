package model

type InvoiceItem struct {
	ID        uint  `gorm:"primaryKey" json:"id"` // ,omitempty
	Quantity  uint  `gorm:"type:smallint; not null" json:"quantity"`
	InvoiceID uint  `gorm:"not null" json:"invoice_id"`
	ProductID uint  `json:"product_id"`
	ServiceID uint8 `json:"service_id"`
	TimeModel
}
