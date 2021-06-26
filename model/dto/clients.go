package dto

import "time"

// optener todos los comentarios
type CommentClients []*CommentClient
type CommentClient struct {
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Comment   string    `gorm:"type:varchar(350); default:''; not null" json:"comment"`
}

type LoginClient struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint8  `json:"role_id"`
}

type SignInClient struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessageInClient struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Affair  string `json:"affair"`
	Message string `json:"message"`
}

// insertar su informacion o actualizar
type InsertClient struct {
	Name     string `gorm:"type:varchar(100); default:''; not null" json:"name"`
	LastName string `gorm:"type:varchar(100); default:''; not null" json:"last_name"`
	Email    string `gorm:"uniqueIndex; type:varchar(255); not null" json:"email"`
	Password string `gorm:"type:varchar(255); not null" json:"password"`
	Dni      string `gorm:"type:char(8)" json:"dni"`
	Ruc      string `gorm:"type:varchar(40); default:''" json:"ruc"`
	Phone    string `gorm:"type:varchar(20); default:''" json:"phone"`
	Picture  string `gorm:"type:varchar(250)" json:"picture"`
	Address  string `gorm:"type:varchar(250); default:''" json:"address"`
}
