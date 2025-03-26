package api

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	// Grupo para a versão 1
	v1 := app.Group("/api/v1")
	RegisterV1Routes(v1)

	// Grupo para a versão 2 (quando necessário)
	// v2 := app.Group("/api/v2")
	// RegisterV2Routes(v2)
}
