package database

import (
	"errors"
	"fmt"
	"github.com/MikelSot/autoPro/model"
)



// IClient interface de CRUD
type IClientCRUD interface {
	Create(client *model.Client) error
	Update(ID uint, client *model.Client) error
	GetByID(ID uint) (*model.Client, error)
	GetAll(Num int) (*model.Clients, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IClient
type IClientQuery interface {
	QueryEmailExists(email string) (bool, error)
	QueryDniExists(dni string) (bool, error)
	//QueryAllAppointment(Num int) (*model.Appointments, error)
	//QueryAllProduct(Num int) (*model.Products, error)
	//QueryOrderAttencionAvailable() ([]string, error)
}


// esta estruvtura ira en el handler
type ClientIN struct {
	iclientCrud IClientCRUD
	iclientQuery IClientQuery
}


// este contructor tambien
func NewClienteIN(iqcrud IClientCRUD, iquery IClientQuery )  ClientIN{
	return ClientIN{
		iclientCrud: iqcrud,
		iclientQuery: iquery,
	}
}


// ClientDAO estructura para hacer referencia a nuestro modelo
type ClientDAO struct {
	clientDao model.Client
}


// NewClient constructor de nuestra estructura, retorna una instancia de  esta
func NewClient() ClientDAO {
	return ClientDAO{}
}


func (c *ClientDAO) Create(client *model.Client) error {
	values := DB().Create(&client)
	fmt.Println("CREAR CLIENTE MENSAJE --> ", values)
	return nil
}


func (c *ClientDAO) Update(ID uint, client *model.Client) error {
	clientID := model.Client{}
	clientID.ID = ID
	values := DB().Model(&clientID).Updates(client)
	fmt.Println("ACTUALIZAR CLIENTE MENSAJE --> ", values)
	return nil
}


func (c *ClientDAO) GetByID(ID uint) (*model.Client, error){
	client := model.Client{}
	values := DB().First(&client, ID)
	//DB().Select("id").First(&client, ID)
	fmt.Println("CONSULTAR CLIENTE POR ID --> ", values)
	fmt.Println("FILAS AFECTADAS POR ID --> ", values.RowsAffected)
	fmt.Println("STATEMENT POR ID --> ", values.Statement)
	return &client, nil
}

func (c *ClientDAO) GetAll(Num int) (*model.Clients, error) {
	clients := model.Clients{}
	values := DB().Limit(Num).Find(&clients)
	fmt.Println("CONSULTAR A TODOS LOS CLIENTES --> ", values)
	return &clients, nil
}


// DeleteSoft borrado sueve, no elimina ese registro como tal de la tabla simplemente le cambia de atributo
func (c *ClientDAO) DeleteSoft(ID uint) error{
	client := model.Client{}
	client.ID = ID
	values := DB().Delete(&client)
	fmt.Println("BORRADO SUAVE --> ", values)
	return nil
}


// DeletePermanent  borrado permanente, borra por completo de la tabla ese registro
func (c *ClientDAO) DeletePermanent(ID uint) error{
	client := model.Client{}
	client.ID = ID
	values := DB().Unscoped().Delete(&client)
	fmt.Println("BORRADO PERMANENTE --> ", values)
	return nil
}

func (c *ClientDAO) QueryEmailExists(email string) (bool, error){
	const  ExistsEmail = "Este Email ya existe"
	client := model.Client{}
	values := DB().Select("Email").Find(&client, "Email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsEmail)
	}
	return false,nil
}

func (c *ClientDAO) QueryDniExists(dni string) (bool, error) {
	const  ExistsDni = "El DNI ya existe"
	client := model.Client{}
	values := DB().Select("Dni").Find(&client, "Dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsDni)
	}
	return false, nil
}













//
//func Crear()  {
//
//	mikel := NewClient()
//	rearME := NewClienteIN(&mikel, &mikel)
//
//	value,err := rearME.iclientQuery.QueryEmailExists("ctmr")
//	if value {
//		fmt.Println("ya esiste el imeal ->", err)
//	}
//}
