package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func Routes(ctx *context.Context, app *fiber.App) {
	// Grupo para a versão 1
	v1 := app.Group("/api/v1")
	RegisterV1Routes(ctx, v1)

	// Grupo para a versão 2 (quando necessário)
	// v2 := app.Group("/api/v2")
	// RegisterV2Routes(v2)
}
