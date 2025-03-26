package api

import (
	_ "api-go-lang/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// New creates a new Fiber instance
func New() *fiber.App {
	app := fiber.New()

	// Registrar as rotas
	RegisterRoutes(app)

	app.Get("/docs/*", swagger.HandlerDefault)

	return app
}
