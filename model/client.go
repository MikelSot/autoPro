package model

type Client struct {
	ID               uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	Name             string `gorm:"type:varchar(100); default:''; not null" json:"name"`
	LastName         string `gorm:"type:varchar(100); default:''; not null" json:"last_name"`
	Email            string `gorm:"uniqueIndex; type:varchar(255); not null" json:"email"`
	Password         string `gorm:"type:varchar(255); not null" json:"password"`
	RememberToken    string `gorm:"type:varchar(250)" json:"remember_token"`
	Dni              string `gorm:"uniqueIndex; type:char(8)" json:"dni"`
	Ruc              string `gorm:"type:varchar(40); default:''" json:"ruc"`
	Phone            string `gorm:"type:varchar(20); default:''" json:"phone"`
	Picture          string `gorm:"type:varchar(250)" json:"picture"`
	Address          string `gorm:"type:varchar(250); default:''" json:"address"`
	State            string `gorm:"type:varchar(25); not null" json:"state"`
	OrderAttention   string `gorm:"type:char(2);default:0; not null" json:"order_attention"`
	Uri              string `gorm:"uniqueIndex; type:varchar(250)" json:"uri"`
	RoleID           uint8  `gorm:"default:2" json:"role_id"`
	Appointments     []Appointment
	TechnicalReviews []TechnicalReview
	Blogs            []Blog
	Comments         []Comment
	Invoices         []Invoice
	TimeModel
}
