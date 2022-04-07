package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnOnce sync.Once
var dbConnection *gorm.DB

func CreateDbConnection() *gorm.DB {

	// only run init of dbConnection once. this should be thred safe
	dbConnOnce.Do(func() {
		var err error

		switch viper.GetString("database.type") {
		case "postgres":
			dbConnection, err = gorm.Open(postgres.New(postgres.Config{
				DSN:                  viper.GetString("database.dsn"),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			}), &gorm.Config{})
		}

		if err != nil {
			fmt.Println(err)
		}
	})

	return dbConnection
}
