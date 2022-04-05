package handler

import (
	"net/http"
	"server/domain/model"

	"github.com/labstack/echo/v4"
)

func GetTask(c echo.Context) error {
	slug := c.Param("slug")

	block := new(model.Block)
	block.ID = 12
	block.Type = "headline"
	block.ParsePayloadFromJsonString(`{
		"submitButtonText": "Abschicken",
		"questions": [
			{
				"text": "Ullamco sint nisi adipisicing Lorem commodo exercitation enim officia culpa tempor?",
				"inputName": "question-1"
			},
			{
				"text": "Culpa consequat ut aliquip id laborum fugiat deserunt quis quis minim veniam?",
				"inputName": "question-2"
			}
		]
	}`)

	task := new(model.Task)
	task.Slug = slug
	task.ID = 123
	task.Blocks = []model.Block{}

	task.AddBlock(*block)

	return c.JSON(http.StatusOK, task)
}
