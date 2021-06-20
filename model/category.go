package model

type Category struct {
	ID     uint8  `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active bool   `gorm:"default:true; not null" json:"active"`
	Offer    bool   `gorm:"default:false; not null" json:"offer"`
	Picture  string `gorm:"type:varchar(250)" json:"picture"`
	Uri      string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	Products []Product
	Blogs    []Blog
	TimeModel
}

type Categories []*Category
