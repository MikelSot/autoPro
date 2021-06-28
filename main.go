package main

import (
	"fmt"
	"regexp"
)

func main() {

	//err := jwt.LoadFiles("jwt/app.rsa", "jwt/app.rsa.public")
	//if err != nil{
	//	log.Fatalf("no se pudo cargar los certificados --> %v",err)
	//	return
	//}
	//
	//database.Migration()
	// a qui hacemos nuestra newclientdao depende


	//e := echo.New()
	// aui traermos las rutas llamamos a esas funciones


	// a qui podemos iniciar el puerto

	err := isEmail("miguelsr.1084@gmail.com")
	if !err {
		fmt.Println("error no es un email")
	}
	fmt.Println("sera un email")
}


func isEmail(email string) bool {
	r, _ := regexp.Compile(`([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`)
	if !r.MatchString(email) {
		return false
	}
	return true
}