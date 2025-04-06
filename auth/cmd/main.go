package main

import (
	"auth/conf"
	"auth/internal/app"
	"context"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg := initConfig()

	app.Run(context.Background(), cfg)
}

func initConfig() *conf.Configuration {
	var config conf.Configuration

	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return &config
}
