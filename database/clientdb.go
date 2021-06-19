package database

import "github.com/MikelSot/autoPro/model"

type IClient interface {
	Create(client *model.Client) error
	Update(ID int, client *model.Client) error
	GetByID(ID int) (*model.Client, error)
	GetAll() (model.Clients, error)
	Delete(ID int)
}


type ClientDAO struct {
	ClientDao model.Client
}

func NewClient(email string, password string) *ClientDAO {
	return &ClientDAO{
		ClientDao:
			model.Client{
				Email: email,
				Password: password,
			},
	}
}

func (c *ClientDAO) Create(client *model.Client)  {
	DB().Create(&client)
}