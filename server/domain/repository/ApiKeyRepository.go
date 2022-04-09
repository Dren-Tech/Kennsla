package repository

import (
	"server/config"
	"server/domain/model"
)

func CreateNewApiKey(key model.ApiKey) uint {
	db := config.CreateDbConnection()

	db.Create(&key)

	return key.ID
}
