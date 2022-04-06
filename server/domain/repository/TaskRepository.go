package repository

import (
	"server/config"
	"server/domain/model"
)

func CreateNewTask(task model.Task) uint {
	return 1337
}

func GetTaskBySlug(slug string) model.Task {
	db := config.CreateDbConnection()
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.Block{})

	var task *model.Task
	db.First(&task, "slug LIKE ?", slug)

	// task.Blocks = []model.Block{}
	// task.AddBlock(*block1)
	// task.AddBlock(*block2)

	return *task
}
