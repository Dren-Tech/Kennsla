package handler

import (
	"net/http"
	"server/domain/repository"

	"github.com/labstack/echo/v4"
)

func GetTask(c echo.Context) error {
	slug := c.Param("slug")

	task := repository.GetTaskBySlug(slug)

	return c.JSON(http.StatusOK, task)
}
