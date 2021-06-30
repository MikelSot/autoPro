package route

import (
	"github.com/MikelSot/autoPro/handler"
	"github.com/labstack/echo/v4"
)

func login(e *echo.Echo,  storage handler.IClientCRUDExists) {
	h := handler.NewLogin(storage)

	e.POST("/v1/login", h.Login)
}
