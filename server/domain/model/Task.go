package model

type Task struct {
	ID     int     `json:"id"`
	Slug   string  `json:"slug"`
	Blocks []Block `json:"blocks"`
}
