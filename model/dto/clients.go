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
type EditClient struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Dni      string `gorm:"type:char(8)" json:"dni"`
	Ruc      string `json:"ruc"`
	Phone    string `json:"phone"`
	Picture  string `json:"picture"`
	Address  string `json:"address"`
}
