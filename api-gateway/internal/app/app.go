package app

import (
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/handler"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Run(cfg *conf.Conf) {
	app := fiber.New()
	
	handler.SetupRoutes(app, cfg)

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
