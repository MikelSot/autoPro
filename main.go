package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
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



	now := time.Now()
	me := now.Add(time.Hour * 24 * 7)
	ahora := now.Format("2006-01-02")
	ahorame := me.Format("2006-01-02")
	//dateString := me.Format("2006-01-02")

	if ahora < ahorame {
		fmt.Println("no son iguales p")
	}

	fmt.Println(ahorame)
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