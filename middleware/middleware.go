package middleware

import (
	"github.com/MikelSot/autoPro/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)


// Authentication para validar que el token valido
func Authentication(fun echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		token := c.Request().Header.Get("Authorization")
		_, err := jwt.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error":"no permitido"})
		}
		return fun(c)
	}
}

// redireccionar si esta autenticado (al home)


// un midlleware para preguntar si es admin