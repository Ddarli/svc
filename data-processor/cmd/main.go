package main

import (
	"context"
	"data-processor/config"
	"data-processor/internal/app"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg := initConfig()

	app.Run(context.Background(), cfg)
}

func initConfig() *config.Config {
	var config config.Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return &config
}
