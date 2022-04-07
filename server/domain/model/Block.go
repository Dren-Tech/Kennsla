package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Type    string         `json:"type"`
	Payload datatypes.JSON `json:"payload"`
	TaskID  uint
}

// func (b *Block) ParsePayloadFromJsonString(payloadString string) {
// 	var result map[string]interface{}

// 	json.Unmarshal([]byte(payloadString), &result)

// 	b.Payload = result
// }

// func (b Block) GetJsonStringFromPayload() string {
// 	str, err := json.Marshal(b.Payload)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return string(str)
// }
