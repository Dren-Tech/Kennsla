package command

import (
	"fmt"
	"server/domain/model"
)

func GenerateNewApiKey() {
	apiKey := model.ApiKey{}
	apiKey.GenerateKey()
	apiKey.IsActive = true

	fmt.Printf("API key is:\n%s\n", apiKey.Key)
}
