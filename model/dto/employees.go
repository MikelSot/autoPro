package dto

import "time"

type AllDataEmployee struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Dni        string    `json:"dni"`
	Ruc        string    `json:"ruc"`
	Phone      string    `json:"phone"`
	Picture    string    `json:"picture"`
	Address    string    `json:"address"`
	State      string    `json:"state"`
	Uri        string    `json:"uri"`
	BirthDate  time.Time `json:"birthdate"`
	Active     bool      `json:"active"`
	Salary     float32   `json:"salary"`
	Turn       string    `json:"turn"`
	Workdays   string    `json:"workdays"`
	Profession string    `json:"profession"`
	BossID     uint      `json:"boss_id"`
	RoleID     uint8     `json:"role_id"`
}


type DataEmployeeHome struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Picture    string    `json:"picture"`
	Uri        string    `json:"uri"`
	Workdays   string    `json:"workdays"`
	Profession string    `json:"profession"`
	RoleID     uint8     `json:"role_id"`
}

type DataEmployeeHomes []*DataEmployeeHome
