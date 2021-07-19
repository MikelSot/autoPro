package model

import "time"

// Employee obtenemos los otros datos del login
type Employee struct {
	ID            uint      `gorm:"primaryKey" json:"id"` // ,omitempty
	Email         string    `gorm:"uniqueIndex; type:varchar(255); not null" json:"email"`
	BirthDate     time.Time `json:"birthdate"`
	Active        bool      `json:"active"`
	Salary        float32   `gorm:"type:real" json:"salary"`
	Turn          string    `gorm:"type:varchar(20)" json:"turn"`
	Workdays      string    `gorm:"type:varchar(255)" json:"workdays"`
	Profession    string    `gorm:"type:varchar(100)" json:"profession"`
	BossID        *uint
	Team          []Employee `gorm:"foreignKey:BossID" json:"boss_id"`
	RoleID        uint8      `gorm:"default:3" json:"role_id"`
	Blogs         []Blog
	Invoices      []Invoice
	TimeModel
}

type Employees []*Employee
