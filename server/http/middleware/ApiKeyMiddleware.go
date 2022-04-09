package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetKeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "TEST", nil
	})
}
