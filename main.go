package main

import (
	"github.com/MikelSot/autoPro/jwt"
	"log"
)

func main() {
	//database.Migration()

	err := jwt.LoadFiles("jwt/app.rsa", "jwt/app.rsa.public")
	if err != nil{
		log.Fatalf("no se pudo cargar los certificados --> %v",err)
		return
	}

}


