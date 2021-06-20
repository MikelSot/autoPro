package database

import (
	"github.com/MikelSot/autoPro/model"
)

type IClient interface {
	Create(client *model.Client) error
	//Create(client *ClientDAO) error
	Update(ID uint, client *model.Client) error
	GetByID(ID uint) (*model.Client, error)
	GetAll(Num int) (*model.Clients, error)
	Delete(ID uint)
}

type ClientIN struct {
	iclient IClient
}

func NewClienteIN(ic IClient)  ClientIN{
	return ClientIN{ic}
}



// queda
type ClientDAO struct {
	clientDao model.Client
}



//func NewClient(email string, password string) *ClientDAO {
//	return &ClientDAO{
//		ClientDao:
//			model.Client{
//				Email: email,
//				Password: password,
//			},
//	}
//}

func NewClient() ClientDAO {
	return ClientDAO{}
}



func (c *ClientDAO) Create(client *model.Client) error {
	DB().Create(&client)
	return nil
}


func (c *ClientDAO) Update(ID uint, client *model.Client) error {
	clientID := model.Client{}
	clientID.ID = ID
	//c.ClientDao.ID = ID
	DB().Model(&clientID).Updates(client)
	return nil
}

func (c *ClientDAO) GetByID(ID uint) (*model.Client, error){
	client := model.Client{}

	DB().First(&client, ID)
	//DB().Select("id").First(&client, ID)
	return &client, nil
}

func (c *ClientDAO) GetAll(Num int) (*model.Clients, error) {
	clients := model.Clients{}
	DB().Limit(Num).Find(&clients)

	return &clients, nil
}


func (c *ClientDAO) Delete(ID uint) {
	panic("implement me")
}

func Crear()  {

	mikel := NewClient()
	rearME := NewClienteIN(&mikel)

	data := model.Client{
		Name: "mikel",
		Email: "mikel@santana",
		Password: "12345",
	}

	rearME.iclient.Create(&data)

}
