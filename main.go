package main

import "github.com/MikelSot/autoPro/database"

//import (
//	"net/http"
//	"github.com/labstack/echo/v4"
//)

func main() {
	database.Migration()

	//me := &model.Client{
	//	Name: "mikel",
	//	Email: "mikel#mikel",
	//	Password: "mikelctmtr",
	//}
	//
	//
	//insertar := database.ClientIN{}
	//insertar.

	//client := database.ClientDAO{}
	//client.Create(&me)


	////client.ClientDao.Name = "Mikel"
	////client.ClientDao.Email = "Mikel@mikel"
	////client.ClientDao.Password = "1233"
	//err := client.Create(&client)

	//if err != nil {
	//	fmt.Println("error al insertar un registro")
	//}

	database.Crear()
}

