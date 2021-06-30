package main

import (
	"fmt"
	"regexp"
	"strings"
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


}

type comer struct {
	name string
}


func isEmail(email string) bool {
	regex :=  `([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`
	r, _ := regexp.Compile(regex)
	if !r.MatchString(email) {
		return false
	}
	return true
}

func com(comer2 *comer) error {
	comer2.name = strings.TrimSpace(comer2.name)
	fmt.Println(comer2.name + "fifi")
	return nil
}