package main

import (
	"context"
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/app"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg := initConfig()

	app.Run(context.Background(), &cfg)
}

func initConfig() conf.Conf {
	var config conf.Conf

	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return config
}
