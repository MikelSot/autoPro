package middleware

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/MikelSot/autoPro/jwt"
)

//  echo.HandlerFunc -->  w http.ResponseWriter, r *http.Request

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