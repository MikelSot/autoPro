package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
)

// IClient interface de CRUD
type IClientCRUD interface {
	Create(clientDot *dto.SignInClient) error
	Update(ID uint, clientDot *dto.InsertClient) error
	GetByID(ID uint) (*model.Client, error)
	GetAll(max int) (*model.Clients, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}


func (c *ClientDao) Create(clientDot *dto.SignInClient) error {
	if len(clientDot.Name) < LenName || len(clientDot.LastName) < LenName {
		return ErrSingIn
	}

	if len(clientDot.Email) < LenName || len(clientDot.Password) < LenName {
		return ErrEmail
	}

	client := model.Client{
		Name:     clientDot.Name,
		LastName: clientDot.LastName,
		Email:    clientDot.Email,
		Password: encrypt(clientDot.Password),
	}
	DB().Create(&client)
	return nil
}

func (c *ClientDao) Update(ID uint, clientDot *dto.InsertClient) error {
	if len(clientDot.Name) < LenName || len(clientDot.LastName) < LenName {
		return ErrUpdate
	}

	if len(clientDot.Email) < LenName || len(clientDot.Password) < LenName {
		return ErrEmail
	}

	client := model.Client{
		Name:     clientDot.Name,
		LastName: clientDot.LastName,
		Email:    clientDot.Email,
		Password: encrypt(clientDot.Password),
		Dni:      clientDot.Dni,
		Ruc:      clientDot.Ruc,
		Phone:    clientDot.Phone,
		Picture:  clientDot.Picture,
		Address:  clientDot.Address,
	}
	clientID := model.Client{}
	clientID.ID = ID
	DB().Model(&clientID).Updates(client)
	return nil
}

func (c *ClientDao) GetByID(ID uint) (*model.Client, error) {
	client := model.Client{}
	DB().First(&client, ID)
	return &client, nil
}

func (c *ClientDao) GetAll(max int) (*model.Clients, error) {
	if max <  MaxGetAll {
		max = MaxGetAll
	}
	clients := model.Clients{}
	DB().Limit(max).Find(&clients)
	return &clients, nil
}

// DeleteSoft borrado sueve, no elimina ese registro como tal de la tabla simplemente le cambia de atributo
func (c *ClientDao) DeleteSoft(ID uint) error {
	client := model.Client{}
	client.ID = ID
	DB().Delete(&client)
	return nil
}

// DeletePermanent  borrado permanente, borra por completo de la tabla ese registro
func (c *ClientDao) DeletePermanent(ID uint) error {
	client := model.Client{}
	client.ID = ID
	DB().Unscoped().Delete(&client)
	return nil
}

func (c *ClientDao) QueryEmailExists(email string) (bool,model.Client, model.Employee, error) {
	const ExistsEmail = "Este Email ya existe USUARIO"
	client := model.Client{}
	values := DB().Limit(1).Select("Email").Find(&client, "email = ?", email)
	//values := DB().Table("clients").Select("Email").Where("email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true, client,model.Employee{},errors.New(ExistsEmail)
	}
	return false,model.Client{}, model.Employee{},nil
}

func (c *ClientDao) QueryDniExists(dni string) (bool, error) {
	const ExistsDni = "El DNI ya existe USUARIO"
	client := model.Client{}
	values := DB().Limit(1).Select("Dni").Find(&client, "dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsDni)
	}
	return false, nil
}

func (c *ClientDao) QueryUriExists(uri string) (bool, error) {
	const ExistsUri = "El DNI ya existe USUARIO"
	client := model.Client{}
	values := DB().Limit(1).Select("Uri").Find(&client, "uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsUri)
	}
	return false, nil
}




