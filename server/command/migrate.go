package command

import (
	"fmt"
	"server/config"
	"server/domain/model"
)

func Migrate() {
	db := config.CreateDbConnection()

	fmt.Println("Starting migration...")

	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.Task{})

	fmt.Println("Finished migration")
}
