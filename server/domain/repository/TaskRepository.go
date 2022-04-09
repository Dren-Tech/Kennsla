package repository

import (
	"fmt"
	"server/config"
	"server/domain/model"

	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

var c *cache.ChainCache = config.GetCache()

func CreateNewTask(task model.Task) uint {
	db := config.CreateDbConnection()

	db.Create(&task)

	return task.ID
}

func GetTaskBySlug(slug string) model.Task {
	db := config.CreateDbConnection()

	var task *model.Task

	cacheResult, _ := c.Get(fmt.Sprintf("task_slug_%s", slug))
	if cacheResult == nil {
		db.Preload("Blocks").First(&task, "slug LIKE ?", slug)
		c.Set(fmt.Sprintf("task_slug_%s", slug), task, &store.Options{})
	} else {
		task = cacheResult.(*model.Task)
	}

	return *task
}
