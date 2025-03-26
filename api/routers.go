package api

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registra as rotas no app Fiber
func RegisterRoutes(app *fiber.App) {
	// @Summary Health Check
	// @Description Verifica se a API está funcionando
	// @Tags Health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /healthcheck [get]
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "API is healthy",
		})
	})

	// @Summary Test endpoint
	// @Description Test endpoint
	// @Tags Test
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /test [get]
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Test endpoint",
		})
	})
}
