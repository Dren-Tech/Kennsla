package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Slug string `json:"slug"`
	//Blocks []Block `json:"blocks"`
}

// func (t *Task) AddBlock(block Block) {
// 	t.Blocks = append(t.Blocks, block)
// }
