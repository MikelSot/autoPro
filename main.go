package main

import (
	"github.com/MikelSot/autoPro/database"
)

//import (
//	"net/http"
//	"github.com/labstack/echo/v4"
//)

func main() {

	if flag := database.FlagMi(); flag != false{
		database.Migration()
	}



}
