package model

type PaymentMethod struct {
	ID       uint8  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active   bool   `gorm:"default:true; not null" json:"active"`
	Invoices []Invoice
	TimeModel
}

type PaymentMethods []*PaymentMethod