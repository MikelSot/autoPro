package model

//Service_Workshops
type Service_Workshops struct {
	ServiceID uint8 `gorm:"not null" json:"Service_id"`
	WorkshopID uint8 `gorm:"not null" json:"Workshop_id"`
}
