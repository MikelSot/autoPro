package model

type Category struct {
	FirstDataModel
	Offer    bool   `gorm:"default:false; not null" json:"offer"`
	Picture  string `gorm:"type:varchar(250)" json:"picture"`
	Uri      string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	Products []Product
	Blogs    []Blog
	TimeModel
}
