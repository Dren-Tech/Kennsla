package repository

import "server/domain/model"

func GetTaskBySlug(slug string) model.Task {
	block1 := new(model.Block)
	block1.ID = 12
	block1.Type = "question"
	block1.ParsePayloadFromJsonString(`{
		"submitButtonText": "Abschicken",
		"questions": [
			{
				"text": "Ullamco sint nisi adipisicing Lorem commodo exercitation enim officia culpa tempor?",
				"inputName": "question-1"
			},
			{
				"text": "Culpa consequat ut aliquip id laborum fugiat deserunt quis quis minim veniam?",
				"inputName": "question-2"
			}
		]
	}`)

	block2 := new(model.Block)
	block2.ID = 821
	block2.Type = "hint"
	block2.ParsePayloadFromJsonString(`{
        "title": "Hinweis Nr. 1",
        "text": "Hinweis-Text 1."
    }`)

	task := new(model.Task)
	task.Slug = slug
	task.ID = 4
	task.Blocks = []model.Block{}

	task.AddBlock(*block1)
	task.AddBlock(*block2)

	return *task
}
