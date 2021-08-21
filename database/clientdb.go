package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"strings"
)

const (
	at = "@"
)

// ClientDAO estructura para hacer referencia a nuestro modelo
type ClientDao struct {
	clientDao model.Client
}

// NewClient constructor de nuestra estructura, retorna una instancia de  esta
func NewClientDao() ClientDao {
	return ClientDao{}
}

func (c *ClientDao) Create(clientDto *dto.SignInClient) error {
	regexSpace := regexp.MustCompile(` `)
	nameWithoutSpace := regexSpace.ReplaceAllString(clientDto.Name, "")
	lastNameWithoutSpace := regexSpace.ReplaceAllString(clientDto.LastName, "")

	url := at + strings.ToLower(nameWithoutSpace) + strings.ToLower(lastNameWithoutSpace)
	existsUri, _ := c.QueryUriExists(url)
	if existsUri {
		url = at + url
	}
	client := model.Client{
		Name:     clientDto.Name,
		LastName: clientDto.LastName,
		Email:    clientDto.Email,
		Password: encrypt(clientDto.Password),
		Uri:      url,
	}
	DB().Create(&client)
	return nil
}

func (c *ClientDao) Update(ID uint, clientDto *dto.EditClient) error {
	client := model.Client{
		Name:     clientDto.Name,
		LastName: clientDto.LastName,
		Email:    clientDto.Email,
		//Password: encrypt(clientDto.Password),
		Dni:      clientDto.Dni,
		Ruc:      clientDto.Ruc,
		Phone:    clientDto.Phone,
		Picture:  clientDto.Picture,
		Address:  clientDto.Address,
		Uri: 	  clientDto.Uri,
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

func (c *ClientDao) UpdatePicture(ID uint, rute string) error {
	DB().Table("clients").Where("id = ?", ID).Update("picture", rute)
	return nil
}

func (c *ClientDao) SelectNameID() (dto.ClientIdNames, error) {
	idNames := dto.ClientIdNames{}
	db.Table("clients").Select("id, name, last_name").Scan(&idNames)
	return idNames,nil
}

func (c *ClientDao) QueryEmailExists(email string) (bool,model.Client, error) {
	client := model.Client{}
	values := DB().Limit(1).Find(&client, "email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true, client,nil
	}
	return false,model.Client{}, nil
}

func (c *ClientDao) QueryDniExists(dni string) (bool,uint,error) {
	client := model.Client{}
	values := DB().Limit(1).Select("Dni", "ID").Find(&client, "dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, client.ID, nil
	}
	return false, uint(0),nil
}

func (c *ClientDao) QueryUriExists(uri string) (bool, error) {
	client := model.Client{}
	values := DB().Limit(1).Select("Uri").Find(&client, "uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, nil
	}
	return false, nil
}


func encrypt(password string) string {
	cost := 6 // es el numero de veces que recorre y encripta
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal("Error al encriptar contraseña\n")
	}
	return string(bytes)
}
