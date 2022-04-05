package handler

import (
	"net/http"
	"server/domain/model"

	"github.com/labstack/echo/v4"
)

func GetTask(c echo.Context) error {
	slug := c.Param("slug")

	task := new(model.Task)
	task.Slug = slug
	task.ID = 123

	return c.JSON(http.StatusOK, task)
}
