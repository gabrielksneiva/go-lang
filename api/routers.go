package api

import (
	"context"
	h "go-lang/api/handlers"

	"github.com/gofiber/fiber/v2"
)

// @Summary Health Check
// @Description Verifica se a API está funcionando
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/healthcheck [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func RegisterV1Routes(ctx *context.Context, group fiber.Router) {
	group.Get("/healthcheck", HealthCheck)
	group.Post("/login", func(c *fiber.Ctx) error {
		return h.LoginHandler(c, ctx)
	})
	group.Post("/register", func(c *fiber.Ctx) error {
		return h.RegisterHandler(c, ctx)
	})
}
