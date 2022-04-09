package model

import (
	"math/rand"

	"gorm.io/gorm"
)

const API_KEY_LENGTH = 24

type ApiKey struct {
	gorm.Model
	Key      string
	IsActive bool
}

func (a *ApiKey) GenerateKey() {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, API_KEY_LENGTH)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	a.Key = string(s)
}
