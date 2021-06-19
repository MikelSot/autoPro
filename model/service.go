package model

type Service struct {
	ID           uint8  `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active       bool   `gorm:"default:true; not null" json:"active"`
	Description  string `gorm:"type:varchar(350); default:''; not null" json:"description"`
	Picture      string `gorm:"type:varchar(250)" json:"picture"`
	Uri          string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	WorkshopID   uint8  `json:"role_id"`
	InvoiceItems []InvoiceItem
	TimeModel
}

type Services []*Service