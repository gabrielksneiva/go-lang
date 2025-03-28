package api

import (
	"context"
	_ "go-lang/docs"
	"go-lang/repositories"
	s "go-lang/services"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @Summary Healthcheck
// @Description Verifica se a aplicação está funcionando corretamente
// @Tags Healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /healthcheck [get]
func Healthcheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "OK"})
}

func NewApp(ctx context.Context) *fiber.App {
	app := fiber.New()
	redisCli, err := repositories.NewRedisClient()
	if err != nil {
		panic(err)
	}

	ls := s.NewLoginService(redisCli)

	SetupRoutes(ctx, app, redisCli, ls)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	app.Get("/healthcheck", Healthcheck)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	return app
}
