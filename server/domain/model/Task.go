package model

type Task struct {
	ID     int     `json:"id"`
	Slug   string  `json:"slug"`
	Blocks []Block `json:"blocks"`
}

func (t *Task) AddBlock(block Block) {
	t.Blocks = append(t.Blocks, block)
}
