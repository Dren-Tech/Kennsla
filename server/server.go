package main

import (
	"fmt"
	"os"
	"server/command"
	"server/config"

	"github.com/eko/gocache/store"
	"github.com/spf13/viper"
)

func main() {
	// load config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// init caches
	cache := config.GetCache()
	cache.Set("cache_health", true, &store.Options{})

	// work on cli arguments
	cli_arguments := os.Args[1:]

	var commandName string
	if len(cli_arguments) >= 1 {
		commandName = cli_arguments[0]
	} else {
		commandName = "start"
	}

	switch commandName {
	case "start":
		command.StartServer()
	case "migrate":
		command.Migrate()
	default:
		command.StartServer()
	}
}
