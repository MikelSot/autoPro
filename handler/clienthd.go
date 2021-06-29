package handler

import (
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	maxLenEmail               = 10
	at                        = "@"
	regexName                 = `([a-zA-Z]{3,})`
	regexDni                  = `([1-9]{8})`
	regexEmail                = `([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`
	ok						  = "ok"
	clientCreated             = "Cliente creado correctamente"
	updatedClient             = "Cliente actualizada correctamente"
	errorEmailIncorrect       = "Email incorrecto"
	errorStruct               = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del correo es invalida"
	errorContent              = "Todos los campos deben tener contenido"
	errorEmailExists          = "Ya existe una cuenta con este correo electrónico"
	errorLenDni               = "Número de identificación inválida"
	errorDniExists            = "Ya existe una cuenta con Número de identificación"
	errorId                   = "Él, id debe ser un número entero positivo"
	errorClientIDDoesNotExist = "El ID del cliente no existe"
	errorGetAll				  = "Hubo un problema al obtener todas los clientes"
)

type clientHd struct {
	crudExists IClientCRUDExists
}

func NewClientHd(ce IClientCRUDExists) clientHd {
	return clientHd{ce}
}

func (c *clientHd) singIn(e echo.Context) error {
	data := dto.SignInClient{}
	err := e.Bind(&data) // de json a estructura
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = areDataValidClient(&data, *c, e)
	if err != nil {
		return err
	}
	err = c.crudExists.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}


	// puede o no ir la generacion de un token
	resp := newResponse(Message, clientCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (c *clientHd) editClient(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := dto.EditClient{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	dataSign := dto.SignInClient{Name: data.Name, LastName: data.LastName, Email: data.Email}
	err = areDataValidClient(&dataSign, *c, e)
	if err != nil {
		return err
	}

	err = isValidDniOrUriClient(&data, *c, e)
	if err != nil {
		return err
	}

	err = c.crudExists.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedClient, nil)
	return e.JSON(http.StatusOK, resp)
	return nil
}

func (c *clientHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := c.crudExists.GetByID(uint(ID))
	if err != nil {
		response := newResponse(Error, errorClientIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (c *clientHd) getAll(e echo.Context) error{
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := c.crudExists.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAll, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (c *clientHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = c.crudExists.DeleteSoft(uint(ID))
	if err != nil {
		response := newResponse(Error, errorClientIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}

func areDataValidClient(data *dto.SignInClient, c clientHd, e echo.Context) error {
	data.Email = strings.TrimSpace(data.Email)
	data.Name = strings.TrimSpace(data.Name)
	data.LastName = strings.TrimSpace(data.LastName)

	if !isEmail(data.Email) {
		resp := newResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if !isEmpty(data.Name) || !isEmpty(data.LastName) || !isEmptyEmail(data.Email) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	exists, _, _, _ := c.crudExists.QueryEmailExists(strings.TrimSpace(data.Email))
	if exists {
		resp := newResponse(Error, errorEmailExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}
	return nil
}

func isValidDniOrUriClient(data *dto.EditClient, c clientHd, e echo.Context) error {
	data.Email = strings.TrimSpace(data.Email)
	data.Name = strings.TrimSpace(data.Name)
	data.LastName = strings.TrimSpace(data.LastName)
	data.Phone = strings.TrimSpace(data.Phone)
	data.Ruc = strings.TrimSpace(data.Ruc)
	data.Address = strings.TrimSpace(data.Address)
	regexSpace := regexp.MustCompile(` `)
	dniWithoutSpace := regexSpace.ReplaceAllString(data.Dni, "")
	nameWithoutSpace := regexSpace.ReplaceAllString(data.Name, "")
	lastNameWithoutSpace := regexSpace.ReplaceAllString(data.LastName, "")

	r, _ := regexp.Compile(regexDni)
	if !r.MatchString(dniWithoutSpace) {
		resp := newResponse(Error, errorLenDni, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	existsDni, _ := c.crudExists.QueryDniExists(dniWithoutSpace)
	if existsDni {
		resp := newResponse(Error, errorDniExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	url := at + strings.ToLower(nameWithoutSpace) + strings.ToLower(lastNameWithoutSpace)
	existsUrl, _ := c.crudExists.QueryUriExists(url)
	if existsUrl {
		url = at + url
	}
	data.Uri = url
	return nil
}

func isEmail(email string) bool {
	r, _ := regexp.Compile(regexEmail)
	if !r.MatchString(email) {
		return false
	}
	return true
}

func isEmpty(string string) bool {
	r, _ := regexp.Compile(regexName)
	if !r.MatchString(string) {
		return false
	}
	return true
}

func isEmptyEmail(string string) bool {
	if len(string) < maxLenEmail {
		return false
	}
	return true
}
