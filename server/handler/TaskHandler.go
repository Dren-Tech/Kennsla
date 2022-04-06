package handler

import (
	"net/http"
	"server/domain/model"
	"server/domain/repository"

	"github.com/labstack/echo/v4"
)

func GetTask(c echo.Context) error {
	slug := c.Param("slug")

	task := repository.GetTaskBySlug(slug)

	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context) error {
	result := repository.CreateNewTask(model.Task{})
	return c.JSON(http.StatusOK, SaveEntityResult{ID: result, Status: "OK"})
}
