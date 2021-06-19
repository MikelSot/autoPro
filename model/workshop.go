package model

type Workshop struct {
	ID       uint8  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active   bool   `gorm:"default:true; not null" json:"active"`
	Address  string `gorm:"type:varchar(250) ; not null" json:"address"`
	Uri      string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	Services []Service
	Products []Product
	Invoices []Invoice
	TimeModel
}

type Workshops []*Workshop
