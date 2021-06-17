package model

type Client struct {
	ID             uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	Name           string `gorm:"type:varchar(100)" json:"name"`
	LastName       string `gorm:"type:varchar(100)" json:"last_name"`
	Email          string `gorm:"type:varchar(255); not null" json:"email"`
	Password       string `gorm:"type:varchar(255); not null" json:"password"`
	RememberToken  string `gorm:"type:varchar(250)" json:"remember_token"`
	Dni            string `gorm:"type:char(8)" json:"dni"`
	Ruc            string `gorm:"type:varchar(40)" json:"ruc"`
	Phone          string `gorm:"type:varchar(20)" json:"phone"`
	Picture        string `gorm:"type:varchar(250)" json:"picture"`
	Address        string `gorm:"type:varchar(250)" json:"address"`
	State          string `gorm:"type:varchar(25)" json:"state"`
	OrderAttention string `gorm:"type:char(2)" json:"order_attention"`
	Uri            string `gorm:"type:varchar(250)" json:"uri"`
	RoleID         uint8  `gorm:"type:varchar(250)" json:"role_id"`
	TimeModel
}
