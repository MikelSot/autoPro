package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	ErrSingIn      = errors.New("error al registrar client")
	ErrUpdate      = errors.New("error al update los datos")
	ErrEmail       = errors.New("error el email o contraseña")
	ErrExistsEmail = errors.New("el email ya existe")
	ErrExistsDni   = errors.New("el email ya existe")
)


// ClientDAO estructura para hacer referencia a nuestro modelo
type ClientDao struct {
	clientDao model.Client
}

// NewClient constructor de nuestra estructura, retorna una instancia de  esta
func NewClientDao() ClientDao {
	return ClientDao{}
}

func (c *ClientDao) Create(client *model.Client) error {
	DB().Create(&client)
	return nil
}

func (c *ClientDao) Update(ID uint, client *model.Client) error {
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
	values := DB().Limit(1).Find(&client, "email = ?", email)
	//values := DB().Limit(1).Select("Email").Find(&client, "email = ?", email)
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


func Encrypt(password string) string {
	cost := 6 // es el numero de veces que recorre y encripta
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal("Error al encriptar contraseña\n")
	}
	return string(bytes)
}
