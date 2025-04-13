package handler

import (
	"github.com/Ddarli/svc/gateway/internal/domain"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const (
	prefix   = "Bearer "
	cookieID = "user_id"
)

func (r *router) jwtProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, prefix) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token missing or invalid",
			})
		}

		token = token[len(prefix):]
		req := domain.ValidateTokenRequest{Token: token}

		res := r.service.ValidateToken(c.Context(), req)

		// TODO: add id hz for what
		if res {
			c.Locals("userID", "TODO")
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token missing or invalid",
			})
		}

		return c.Next()
	}
}

func (r *router) cookieAuthMiddleware(ctx *fiber.Ctx) error {
	userID := ctx.Cookies(cookieID)
	if userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID cookie missing",
		})
	}

	ctx.Locals(cookieID, userID)

	return ctx.Next()
}
