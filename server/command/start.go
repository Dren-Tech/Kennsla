package command

import (
	"net/http"
	"server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()

	// MIDDLEWARE

	// log requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}] ${uri} (${status}) [${latency_human}]\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// task handlers
	e.GET("/task/:slug", handler.GetTask)
	e.POST("/task", handler.CreateTask)

	e.Logger.Fatal(e.Start(":1323"))
}
