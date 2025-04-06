package router

import (
	"auth/conf"
	"auth/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app     *fiber.App
	service service.Service
}

func NewRouter(app *fiber.App, service service.Service) *Router {
	return &Router{app: app, service: service}
}

func (r *Router) RegisterRoutes(cfg *conf.Configuration) {
	r.app.Post(cfg.Routes.Register, r.handleRegister)
	r.app.Post(cfg.Routes.Authenticate, r.handleAuth)
}
