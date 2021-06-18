package model

type Service struct {
	FirstDataModel
	Description  string `gorm:"type:varchar(350); default:''; not null" json:"description"`
	Picture      string `gorm:"type:varchar(250)" json:"picture"`
	Uri          string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	WorkshopID   uint8  `json:"role_id"`
	InvoiceItems []InvoiceItem
	TimeModel
}
