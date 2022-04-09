package repository

import (
	"fmt"
	"server/config"
	"server/domain/model"

	"github.com/eko/gocache/store"
)

func CreateNewApiKey(key model.ApiKey) uint {
	db := config.CreateDbConnection()

	db.Create(&key)

	return key.ID
}

func GetApiKey(key string) model.ApiKey {
	db := config.CreateDbConnection()

	var apiKey *model.ApiKey

	cacheResult, _ := c.Get(fmt.Sprintf("apikey_key_%s", key))
	if cacheResult == nil {
		fmt.Println("KEY IS NOT IN CACHE")
		db.First(&apiKey, "key LIKE ? AND is_active IS true", key)
		c.Set(fmt.Sprintf("apikey_key_%s", key), apiKey, &store.Options{})
	} else {
		apiKey = cacheResult.(*model.ApiKey)
	}

	return *apiKey
}
