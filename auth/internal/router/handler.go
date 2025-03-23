package router

import (
	"auth/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) handleAuth(ctx *fiber.Ctx) error {
	var req model.AuthRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(200)
}

func (r *Router) handleRegister(ctx *fiber.Ctx) error {
	var req model.RegisterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := r.service.NewUser(ctx.Context(), req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(200)
}
