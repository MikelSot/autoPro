package handler

import (
	"github.com/labstack/echo"
)

const (
	errorStructPicture = "Error en la estructura al subir una imagen"
)


func UploadFile(e echo.Context) error{

	return nil
}

//  serve los archivos estaticos
//func main() {
//	e := echo.New()
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//	e.Static("/", "public")
//	e.POST("/upload/:id", UploadFile)
//	e.Logger.Fatal(e.Start(":1323"))
//}
