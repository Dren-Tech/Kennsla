package handler

import (
	"encoding/json"
	"fmt"
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
	var task model.Task

	// get json from body and decode it into task
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		panic(fmt.Errorf("fatal error: %w", err))
	}

	result := repository.CreateNewTask(task)
	return c.JSON(http.StatusOK, SaveEntityResult{ID: result, Status: "OK"})
}
