package route

import (
	"github.com/MikelSot/autoPro/database"
	"github.com/MikelSot/autoPro/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func StartServer()  {
	err := jwt.LoadFiles("jwt/app.rsa", "jwt/app.rsa.public")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	database.Migration()

	// a qui hacemos nuestra newclientdao depende

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// aui traermos las rutas llamamos a esas funciones

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
