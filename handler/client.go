package handler

import (
	"encoding/json"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo"
	"net/http"
	"regexp"
)

const (
	maxLenName           = 3
	maxLenEmail          = 10
	clientCreated		= "Cliente creado correctamente"
	errorEmailIncorrect  = "Email incorrecto"
	errorEmailOrPassword = "El email o contraseña son incorrectos"
	errorStruct          = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del correo es invalida"
	errorContent 		 = "Todos los campos deben tener contenido"
	errorEmailExists 	 = "Ya existe una cuenta con este correo electrónico"
)

type client struct {
	query IQueryExists
	crud  IClientCRUD
}

func NewClient(q IQueryExists, c IClientCRUD) client {
	return client{
		q,
		c,
	}
}


func (c *client) singIn(e echo.Context) error {
	data := dto.SignInClient{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = areDataValid(data, *c, e)
	if err != nil {
		return err
	}

	err = c.crud.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, clientCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (c *client) editClient(w http.ResponseWriter, r *http.Request) {

	data := dto.EditClient{}
	err := json.NewDecoder(r.Body).Decode(&data) // de json a estructura
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	err = c.crud.Update(uint(4), &data)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":""message", "message":"correcto"}`))
}

func (c *client) getById(w http.ResponseWriter, r *http.Request) {

}

func (c *client) getAll(w http.ResponseWriter, r *http.Request) {

}

func areDataValid(data dto.SignInClient,c client ,e echo.Context) error {
	if !isEmail(data.Email) {
		resp := newResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if !isEmpty(data.Name) || !isEmpty(data.LastName) || !isEmptyEmail(data.Email) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	exists, _, _, _ := c.query.QueryEmailExists(data.Email)
	if exists {
		resp := newResponse(Error, errorEmailExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}
	return nil
}

func isEmail(email string) bool {
	r, _ := regexp.Compile(`([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`)
	if !r.MatchString(email) {
		return false
	}
	return true
}

func isEmpty(string string) bool {
	if len(string) < maxLenName {
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