package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	ErrSingIn = errors.New("error al registrar client")
	ErrUpdate = errors.New("error al update los datos")
	ErrEmail  = errors.New("error el email o contraseña")
)




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


func (c *ClientDao) Login(email string, password string) (model.Client, bool) {
	exists,client,_,_:= c.QueryEmailExists(email)
	if exists == false{
		return client, false
	}

	pass := []byte(password)
	passdb := []byte(client.Password)
	err := bcrypt.CompareHashAndPassword(passdb, pass)
	if err != nil{
		// el error ocurre cuando no coinciden
		return model.Client{}, false
	}
	return client,true
}


func encrypt(password string) string{
	cost := 6 // es el numero de veces que recorre y encripta
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal("Error al encriptar contraseña\n")
	}
	return string(bytes)
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
