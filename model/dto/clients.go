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
	UpdatedAt time.Time `json:"updated_at"`
	Comment   string    `json:"comment"`
}

// iniciar sesion
type LoginClient struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// registrase
type SignInClient struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//actualizar informacion
type EditClient struct {
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

// datos del cliente al iniciar sesion
type DataClient struct {
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
	State    string `json:"state"`
	Uri      string `json:"uri"`
	Role     uint8  `json:"role"`
}

// Claim contiene los datos que iran en el payload del token
type Claim struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Uri      string `json:"uri"`
	State    string `json:"state"`
	Role     uint8  `json:"role"`
	jwt.StandardClaims
}

type UploadAvatar struct {
	ID       uint   `json:"id"`
	Picture  string `json:"picture"`
}