package dto

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
	Uri      string `json:"uri"`
}

type DataClient struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	RememberToken string `json:"remember_token"`
	Dni           string `gorm:"type:char(8)" json:"dni"`
	Ruc           string `json:"ruc"`
	Phone         string `json:"phone"`
	Picture       string `json:"picture"`
	Address       string `json:"address"`
	State         string `json:"state"`
	Role          uint8  `json:"role"`
}

// Claim contiene los datos que iran en el payload del token
type Claim struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	State    string `json:"state"`
	Role     uint8  `json:"role"`
	jwt.StandardClaims
}
