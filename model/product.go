package model

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(200); not null" json:"name"`
	SKU         string  `gorm:"type:varchar(50); default:''" json:"sku"`
	ProductCode string  `gorm:"type:varchar(100)" json:"product_code"`
	Measures    string  `gorm:"type:varchar(50)" json:"measure"`
	Active      bool    `gorm:"default:false; not null" json:"active"`
	Stock       uint    `gorm:"type:smallint; not null" json:"stock"`
	UnitPrice   float32 `gorm:"type:real; not null" json:"unit_price"`
	Uri         string  `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	Offer       bool    `gorm:"default:false; not null" json:"offer"`
	Picture     string  `gorm:"type:varchar(250)" json:"picture"`
	Description string  `gorm:"type:varchar(350); default:''; not null" json:"description"`
	CategoryID  uint8   `json:"category_id"`
	WorkshopID  uint8   `json:"workshop_id"`
	Comments    []Comment
	TimeModel
}
