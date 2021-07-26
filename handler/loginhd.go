package handler

import (
	"github.com/MikelSot/autoPro/jwt"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

const (
	errorGenerateToken   = "No se pudo generar el token"
	errorEmailOrPassword = "El email o contraseña son incorrectos"
)

type login struct {
	crudExists IClientCRUDExists
}

func NewLogin(ce IClientCRUDExists) login {
	return login{ce}
}

func (l *login) Login(e echo.Context) error {
	data := dto.LoginClient{}
	err := e.Bind(&data) // de json a estructura
	if err != nil {
		resp := NewResponse(Error, errorEmailOrPassword, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Email = strings.TrimSpace(data.Email)
	if !isEmail(data.Email) {
		resp := NewResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	exists, client, _ := l.crudExists.QueryEmailExists(data.Email)
	if !exists {
		resp := NewResponse(Error, errorEmailIncorrect, client)
		return e.JSON(http.StatusBadRequest, resp)
	}

	pass := []byte(data.Password)
	passDB := []byte(client.Password)
	err = bcrypt.CompareHashAndPassword(passDB, pass)
	if err != nil {
		resp := NewResponse(Error, errorEmailOrPassword, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}
	password = data.Password // getByID

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
		Uri:      client.Uri,
		Role:     client.RoleID,
	}
	token, err := jwt.GenerateToken(&dataClient)
	if err != nil {
		resp := NewResponse(Error, errorGenerateToken, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, ok, token)
	return e.JSON(http.StatusOK, resp)
}
