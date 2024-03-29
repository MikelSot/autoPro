package model

type Client struct {
	ID               uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	Name             string `gorm:"type:varchar(100); default:''; not null" json:"name"`
	LastName         string `gorm:"type:varchar(100); default:''; not null" json:"last_name"`
	Email            string `gorm:"uniqueIndex; type:varchar(255); not null" json:"email"`
	Password         string `gorm:"type:varchar(255); not null" json:"password"`
	Dni              string `gorm:"type:char(8)" json:"dni"`
	Ruc              string `gorm:"type:varchar(40); default:''" json:"ruc"`
	Phone            string `gorm:"type:varchar(20); default:''" json:"phone"`
	Picture          string `gorm:"type:varchar(250)" json:"picture"`
	Address          string `gorm:"type:varchar(250); default:''" json:"address"`
	State            string `gorm:"type:varchar(25); default:'active'; not null" json:"state"`
	Uri              string `gorm:"type:varchar(250)" json:"uri"`
	RoleID           uint8  `gorm:"default:2" json:"role_id"`
	Appointments     []Appointment
	TechnicalReviews []TechnicalReview
	Comments         []Comment
	Invoices         []Invoice
	TimeModel
}
type Clients []*Client
