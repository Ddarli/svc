package app

import (
	"context"
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/client"
	"github.com/Ddarli/svc/gateway/internal/handler"
	"github.com/Ddarli/svc/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func Run(ctx context.Context, cfg *conf.Conf) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	authClient := client.New(ctx, &cfg.Client)
	dataClient := client.NewDataClient(ctx, &cfg.Client)

	service := service.New(authClient, dataClient)
	router := handler.New(app, cfg, service)

	router.SetupRoutes()

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
