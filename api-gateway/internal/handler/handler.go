package handler

import (
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/domain"
	"github.com/Ddarli/svc/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	app     *fiber.App
	cfg     *conf.Conf
	service *service.Service
}

func New(app *fiber.App, cfg *conf.Conf, service *service.Service) *router {
	return &router{
		app:     app,
		cfg:     cfg,
		service: service,
	}
}

func (r *router) SetupRoutes() {
	r.app.Post("/auth/authenticate", r.handleAuth)
	r.app.Post("/auth/register", r.handleRegistration)
	r.app.Get("api/v1/blabla", r.jwtProtected(), r.handleBlabla)
}

func (r *router) handleAuth(ctx *fiber.Ctx) error {
	var loginRequest domain.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	resp := r.service.Auth(ctx.Context(), loginRequest)
	if resp == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Token": resp,
	})
}

func (r *router) handleRegistration(ctx *fiber.Ctx) error {
	var registerRequest domain.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	resp := r.service.Register(ctx.Context(), registerRequest)
	if resp == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Token": resp,
	})
}

func (r *router) handleBlabla(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"answer": "blablalba",
	})
}
