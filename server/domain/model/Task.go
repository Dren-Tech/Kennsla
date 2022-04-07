package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title  string  `json:"title"`
	Slug   string  `json:"slug"`
	Blocks []Block `json:"blocks"`
}

func (t *Task) AddBlock(block Block) {
	t.Blocks = append(t.Blocks, block)
}

func (t *Task) SluggifyTitle() string {
	return slug.Make(t.Title)
}
