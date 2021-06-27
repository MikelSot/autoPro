package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
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

func (c *ClientDao) Login(clientdto dto.LoginClient) (model.Client, bool) {
	exists, client, _, _ := c.QueryEmailExists(clientdto.Email)
	if exists == false {
		return client, false
	}

	pass := []byte(clientdto.Password)
	passdb := []byte(client.Password)
	err := bcrypt.CompareHashAndPassword(passdb, pass)
	if err != nil {
		// el error ocurre cuando no coinciden
		return model.Client{}, false
	}
	return client, true
}

func (c *ClientDao) SingIn(clientDto dto.SignInClient) (bool, error) {
	if len(clientDto.Name) < LenName || len(clientDto.LastName) < LenName {
		return false, ErrSingIn
	}

	if len(clientDto.Email) < LenName || len(clientDto.Password) < LenName {
		return false, ErrEmail
	}

	existsEmail, _, _, _ := c.QueryEmailExists(clientDto.Email)
	if existsEmail {
		return false, ErrExistsEmail
	}

	url := at + clientDto.Name + clientDto.LastName
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

	err := c.Create(&client)
	if err != nil {
		return false, ErrSingIn
	}
	return true, nil
}

func (c *ClientDao) EditInformation(clientDto dto.EditClient) (bool, error) {
	if len(clientDto.Name) < LenName || len(clientDto.LastName) < LenName {
		return false, ErrUpdate
	}

	if len(clientDto.Email) < LenName || len(clientDto.Password) < LenName {
		return false, ErrEmail
	}

	existsEmail, _, _, _ := c.QueryEmailExists(clientDto.Email)
	if existsEmail {
		return false, ErrExistsEmail
	}

	existsDni, _ := c.QueryDniExists(clientDto.Dni)
	if existsDni {
		return false, ErrExistsDni
	}

	client := model.Client{
		Name:     clientDto.Name,
		LastName: clientDto.LastName,
		Email:    clientDto.Email,
		Password: encrypt(clientDto.Password),
		Dni:      clientDto.Dni,
		Ruc:      clientDto.Ruc,
		Phone:    clientDto.Phone,
		Picture:  clientDto.Picture,
		Address:  clientDto.Address,
	}
	err := c.Update(clientDto.ID, &client)
	if err != nil {
		return false, ErrSingIn
	}
	return true, nil
}

func encrypt(password string) string {
	cost := 6 // es el numero de veces que recorre y encripta
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal("Error al encriptar contraseña\n")
	}
	return string(bytes)
}
