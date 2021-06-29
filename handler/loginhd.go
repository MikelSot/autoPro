package handler

import (
	"github.com/MikelSot/autoPro/jwt"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

//func (c *ClientDao) Login(clientdto dto.LoginClient) (model.Client, bool) {
//	exists, client, _, _ := c.QueryEmailExists(clientdto.Email)
//	if exists == false {
//		return client, false
//	}
//
//	pass := []byte(clientdto.Password)
//	passdb := []byte(client.Password)
//	err := bcrypt.CompareHashAndPassword(passdb, pass)
//	if err != nil {
//		// el error ocurre cuando no coinciden
//		return model.Client{}, false
//	}
//	return client, true
//}
//

//func (c *ClientDao) Create(client *model.Client) error {
//	DB().Create(&client)
//	return nil
//}
//
//func (c *ClientDao) Update(ID uint, client *model.Client) error {
//	clientID := model.Client{}
//	clientID.ID = ID
//	DB().Model(&clientID).Updates(client)
//	return nil
//}

const (
	errorGenerateToken   = "No se pudo generar el token"
	errorEmailOrPassword = "El email o contraseña son incorrectos"
)

type login struct {
	crudExists IClientCRUDExists
}

func newLogin(ce IClientCRUDExists) login {
	return login{ce}
}

func (l *login) login(e echo.Context) error {
	data := dto.LoginClient{}
	err := e.Bind(&data) // de json a estructura
	if err != nil {
		resp := newResponse(Error, errorEmailOrPassword, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Email = strings.TrimSpace(data.Email)
	if !isEmail(data.Email) {
		resp := newResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	exists, client, _, _ := l.crudExists.QueryEmailExists(data.Email)
	if !exists {
		resp := newResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	pass := []byte(data.Password)
	passDB := []byte(client.Password)
	err = bcrypt.CompareHashAndPassword(passDB, pass)
	if err != nil {
		resp := newResponse(Error, errorEmailOrPassword, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	dataClient := dto.DataClient{
		ID:       client.ID,
		Name:     client.Name,
		LastName: client.LastName,
		Email:    client.Email,
		Password: client.Password,
		Dni:      client.Dni,
		Ruc:      client.Ruc,
		Phone:    client.Phone,
		Picture:  client.Picture,
		Address:  client.Address,
		State:    client.State,
		Role:     client.RoleID,
	}
	token, err := jwt.GenerateToken(&dataClient)
	if err != nil {
		resp := newResponse(Error, errorGenerateToken, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, ok, token)
	return e.JSON(http.StatusOK, resp)
}
