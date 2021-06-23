package model

type Comment struct {
	ID            uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	ClientID      uint   `json:"client_id"`
	ProductID     uint   `gorm:"default:5" json:"product_id"`
	Qualification uint8  `json:"qualification"`
	Comment       string `gorm:"type:varchar(350); default:''; not null" json:"comment"`
	BlogID        uint   `json:"blog_id"`
	TimeModel
}

type Comments []*Comment