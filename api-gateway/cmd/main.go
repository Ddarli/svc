package main

import (
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg := initConfig()

	app := fiber.New()

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}

func initConfig() conf.Conf {
	var config conf.Conf

	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	return config
}
