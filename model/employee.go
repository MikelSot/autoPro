package model

import "time"

type Employee struct {
	ID            uint      `gorm:"primaryKey" json:"id"` // ,omitempty
	Name          string    `gorm:"type:varchar(100); default:''; not null" json:"name"`
	LastName      string    `gorm:"type:varchar(100); default:''; not null" json:"last_name"`
	Email         string    `gorm:"uniqueIndex; type:varchar(255); not null" json:"email"`
	Password      string    `gorm:"type:varchar(255); not null" json:"password"`
	RememberToken string    `gorm:"type:varchar(250)" json:"remember_token"`
	Dni           string    `gorm:"uniqueIndex; type:char(8)" json:"dni"`
	BirthDate     time.Time `json:"birthdate"`
	Phone         string    `gorm:"type:varchar(20); default:''" json:"phone"`
	Picture       string    `gorm:"type:varchar(250)" json:"picture"`
	Address       string    `gorm:"type:varchar(250); default:''" json:"address"`
	Active        bool      `json:"active"`
	Salary        float32   `gorm:"type:real" json:"salary"`
	Turn          string    `gorm:"type:varchar(20)" json:"turn"`
	Workdays      string    `gorm:"type:varchar(255)" json:"workdays"`
	Profession    string    `gorm:"type:varchar(100)" json:"profession"`
	Uri           string    `gorm:"type:varchar(250)" json:"uri"`
	Description   string    `gorm:"type:varchar(350); default:''" json:"description"`
	BossID        *uint
	Team          []Employee `gorm:"foreignKey:BossID" json:"boss_id"`
	RoleID        uint8      `json:"role_id"`
	Appointments  []Appointment
	TimeModel
}
