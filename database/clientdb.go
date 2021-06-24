package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
)



// IClient interface de CRUD
type IClientCRUD interface {
	Create(client *model.Client) error
	Update(ID uint, client *model.Client) error
	GetByID(ID uint) (*model.Client, error)
	GetAll(max int) (*model.Clients, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}



//// esta estruvtura ira en el handler
//type ClientIN struct {
//	iclientCrud IClientCRUD
//	iclientQuery IQueryExists
//}
//
//
//// este contructor tambien
//func NewClienteIN(iqcrud IClientCRUD, iexists IQueryExists )  ClientIN{
//	return ClientIN{
//		iclientCrud: iqcrud,
//		iclientQuery: iexists,
//	}
//}


// ClientDAO estructura para hacer referencia a nuestro modelo
type ClientDao struct {
	clientDao model.Client
}


// NewClient constructor de nuestra estructura, retorna una instancia de  esta
func NewClientDao() ClientDao {
	return ClientDao{}
}


func (c *ClientDao) Create(client *model.Client) error {
	var err error = errors.New("Ingresar el campo nombre o apellido")
	if len(client.Name) < LenName || len(client.LastName) < LenName{
		return err
	}
	DB().Create(&client)
	return nil
}


func (c *ClientDao) Update(ID uint, client *model.Client) error {
	var err error = errors.New("Ingresar el campo nombre o apellido")
	if len(client.Name) < LenName || len(client.LastName) < LenName{
		return err
	}
	clientID := model.Client{}
	clientID.ID = ID
	DB().Model(&clientID).Updates(client)
	return nil
}


func (c *ClientDao) GetByID(ID uint) (*model.Client, error){
	client := model.Client{}
	DB().First(&client, ID)
	return &client, nil
}

func (c *ClientDao) GetAll(max int) (*model.Clients, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	clients := model.Clients{}
	DB().Limit(max).Find(&clients)
	return &clients, nil
}


// DeleteSoft borrado sueve, no elimina ese registro como tal de la tabla simplemente le cambia de atributo
func (c *ClientDao) DeleteSoft(ID uint) error{
	client := model.Client{}
	client.ID = ID
	DB().Delete(&client)
	return nil
}


// DeletePermanent  borrado permanente, borra por completo de la tabla ese registro
func (c *ClientDao) DeletePermanent(ID uint) error{
	client := model.Client{}
	client.ID = ID
	DB().Unscoped().Delete(&client)
	return nil
}

func (c *ClientDao) QueryEmailExists(email string) (bool, error){
	const  ExistsEmail = "Este Email ya existe USUARIO"
	client := model.Client{}
	values := DB().Select("Email").Find(&client, "email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsEmail)
	}
	return false,nil
}

func (c *ClientDao) QueryDniExists(dni string) (bool, error) {
	const  ExistsDni = "El DNI ya existe USUARIO"
	client := model.Client{}
	values := DB().Select("Dni").Find(&client, "dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsDni)
	}
	return false, nil
}

func (c *ClientDao)	QueryUriExists(uri string) (bool, error) {
	const  ExistsUri = "El DNI ya existe USUARIO"
	client := model.Client{}
	values := DB().Select("Uri").Find(&client, "uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsUri)
	}
	return false,nil
}












//func Crear()  {
//
//	mikel := NewClient()
//	rearME := NewClienteIN(&mikel, &mikel)
//
//	value,err := rearME.iclientQuery.QueryEmailExists("ctmr")
//	if value {
//		fmt.Printf("ya esiste el imeal -> %v", err)
//	}
//}
