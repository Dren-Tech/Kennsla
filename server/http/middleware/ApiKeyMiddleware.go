package middleware

import (
	"os"
	"server/domain/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetKeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		if os.Getenv("APP_ENV") == "dev" {
			return true, nil
		}
		apiKey := repository.GetApiKey(key)
		return apiKey.Key != "", nil
	})
}
