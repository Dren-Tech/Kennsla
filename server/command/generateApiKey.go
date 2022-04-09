package command

import (
	"fmt"
	"server/domain/model"
	"server/domain/repository"
)

func GenerateNewApiKey() {
	apiKey := model.ApiKey{}
	apiKey.GenerateKey()
	apiKey.IsActive = true

	repository.CreateNewApiKey(apiKey)

	fmt.Printf("API key is:\n%s\n", apiKey.Key)
}
