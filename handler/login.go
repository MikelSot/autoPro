package handler

import (
	"fmt"
	"github.com/MikelSot/autoPro/jwt"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo"
	"net/http"
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
//func (c *ClientDao) SingIn(clientDto dto.SignInClient) (bool, error) {
//	if len(clientDto.Name) < LenName || len(clientDto.LastName) < LenName {
//		return false, ErrSingIn
//	}
//
//	if len(clientDto.Email) < LenName || len(clientDto.Password) < LenName {
//		return false, ErrEmail
//	}
//
//	existsEmail, _, _, _ := c.QueryEmailExists(clientDto.Email)
//	if existsEmail {
//		return false, ErrExistsEmail
//	}
//
//	url := at + clientDto.Name + clientDto.LastName
//	existsUri, _ := c.QueryUriExists(url)
//	if existsUri {
//		url = at + url
//	}
//	client := model.Client{
//		Name:     clientDto.Name,
//		LastName: clientDto.LastName,
//		Email:    clientDto.Email,
//		Password: Encrypt(clientDto.Password),
//		Uri:      url,
//	}
//
//	err := c.Create(&client)
//	if err != nil {
//		return false, ErrSingIn
//	}
//	return true, nil
//}
//
//func (c *ClientDao) EditInformation(clientDto dto.EditClient) (bool, error) {
//	if len(clientDto.Name) < LenName || len(clientDto.LastName) < LenName {
//		return false, ErrUpdate
//	}
//
//	if len(clientDto.Email) < LenName || len(clientDto.Password) < LenName {
//		return false, ErrEmail
//	}
//
//	existsEmail, _, _, _ := c.QueryEmailExists(clientDto.Email)
//	if existsEmail {
//		return false, ErrExistsEmail
//	}
//
//	existsDni, _ := c.QueryDniExists(clientDto.Dni)
//	if existsDni {
//		return false, ErrExistsDni
//	}
//
//	client := model.Client{
//		Name:     clientDto.Name,
//		LastName: clientDto.LastName,
//		Email:    clientDto.Email,
//		Password: Encrypt(clientDto.Password),
//		Dni:      clientDto.Dni,
//		Ruc:      clientDto.Ruc,
//		Phone:    clientDto.Phone,
//		Picture:  clientDto.Picture,
//		Address:  clientDto.Address,
//	}
//	err := c.Update(clientDto.ID, &client)
//	if err != nil {
//		return false, ErrSingIn
//	}
//	return true, nil
//}




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

type login struct {
	query IQueryExists
	crud IClientCRUD
}

func newLogin(q IQueryExists, c IClientCRUD) login {
	return login{
		q,
		c,
	}
}

//
//

// Todos los campos deben tener contenido

func (l *login) login(c echo.Context) error {
	// validar con regex

	data :=dto.LoginClient{}
	err := c.Bind(&data)  //err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		resp := newResponse(Error, "estructura no valida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	// a qui validas si existe el email y contraseña son validos y si existen

	// ya luego de validar eso generas el token (recuerda que puedes retornar los datos de ese cliente)
	token, err := jwt.GenerateToken(&data)
	if err != nil{
		resp := newResponse(Error, "estructura no valida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	fmt.Println(token)
	// si salio bien all le enviamos la data  que queremos enviar
	return nil

}




