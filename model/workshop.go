package model

type Workshop struct {
	FirstDataModel
	Address string `gorm:"type:varchar(250) ; not null" json:"address"`
	Uri     string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	Services []Service
	TimeModel
}

