package api

import (
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

func RegisterV1Routes(group fiber.Router) {
	group.Get("/healthcheck", HealthCheck)

}
