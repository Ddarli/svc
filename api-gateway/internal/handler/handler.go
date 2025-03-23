package handler

import (
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func SetupRoutes(app *fiber.App, cfg *conf.Conf) {
	app.Get("/auth/register", proxyHandler(cfg.Services.AuthService+"/register"))
	app.Get("/auth/authenticate", proxyHandler(cfg.Services.AuthService+"/authenticate"))
}

func proxyHandler(target string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		url := target + ctx.OriginalURL()

		return proxy.Do(ctx, url)
	}
}
