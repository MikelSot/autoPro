package model

type Service struct {
	ID           uint8      `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active       bool       `gorm:"default:true; not null" json:"active"`
	Description  string     `gorm:"type:varchar(350); default:''; not null" json:"description"`
	Picture      string     `gorm:"type:varchar(250)" json:"picture"`
	Uri          string     `gorm:"type:varchar(250); default:services.name" json:"uri"`
	Workshops    []Workshop `gorm:"many2many:service_workshops;" json:"workshops"` // muchos a muchos
	InvoiceItems []InvoiceItem
	TimeModel
}

type Services []*Service
