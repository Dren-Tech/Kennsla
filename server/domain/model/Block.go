package model

import (
	"encoding/json"
	"fmt"
)

type Block struct {
	ID      int                    `json:"id"`
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

func (b *Block) ParsePayloadFromJsonString(payloadString string) {
	var result map[string]interface{}

	json.Unmarshal([]byte(payloadString), &result)

	b.Payload = result
}

func (b Block) GetJsonStringFromPayload() string {
	str, err := json.Marshal(b.Payload)

	if err != nil {
		fmt.Println(err)
	}

	return string(str)
}
