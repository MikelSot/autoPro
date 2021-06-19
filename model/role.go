package model

type Role struct {
	ID     uint8  `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"uniqueIndex; type:varchar(50) ; not null" json:"name"`
	Active bool   `gorm:"default:true; not null" json:"active"`
	Employees []Employee
	Clients   []Client
	TimeModel
}

type Roles []*Role
