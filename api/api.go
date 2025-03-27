package api

import (
	"context"
	_ "go-lang/docs"
	r "go-lang/repositories"

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

// NewApp cria e configura uma nova instância do Fiber
func NewApp(ctx context.Context, r r.RedisClient) *fiber.App {
	// Inicializa o Fiber
	app := fiber.New()

	// Configura as rotas
	SetupRoutes(ctx, app, r)

	// Configura o Swagger
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// Configura a rota de healthcheck
	app.Get("/healthcheck", Healthcheck)

	// Configura o Swagger
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	return app
}
