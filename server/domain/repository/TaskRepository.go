package repository

import (
	"server/config"
	"server/domain/model"
)

func CreateNewTask(task model.Task) uint {
	db := config.CreateDbConnection()

	db.Create(&task)

	return task.ID
}

func GetTaskBySlug(slug string) model.Task {
	db := config.CreateDbConnection()
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.Block{})

	var task *model.Task
	db.Preload("Blocks").First(&task, "slug LIKE ?", slug)

	return *task
}
