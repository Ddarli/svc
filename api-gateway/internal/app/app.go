package app

import (
	"context"
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/client"
	"github.com/Ddarli/svc/gateway/internal/handler"
	"github.com/Ddarli/svc/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Run(ctx context.Context, cfg *conf.Conf) {
	app := fiber.New()

	authClient := client.New(ctx, &cfg.Client)
	service := service.New(authClient)
	router := handler.New(app, cfg, service)

	router.SetupRoutes()

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
