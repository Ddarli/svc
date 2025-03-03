package app

import (
	"auth/conf"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Run(configuration *conf.Configuration) {
	app := fiber.New()

	log.Fatal(app.Listen(":" + configuration.Server.Port))
}
