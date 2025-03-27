package api

import (
	"context"
	h "go-lang/api/handlers"
	r "go-lang/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(ctx context.Context, app *fiber.App, r r.RedisClient) {
	v1 := app.Group("/api/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {
		return h.LoginHandler(c, ctx, r)
	})
	v1.Post("/register", func(c *fiber.Ctx) error {
		return h.RegisterHandler(c, ctx, r)
	})

}
