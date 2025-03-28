package api

import (
	"context"
	h "go-lang/api/handlers"
	repositories "go-lang/repositories"
	s "go-lang/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(ctx context.Context, app *fiber.App, r *repositories.RedisClient, ls *s.LoginService) {
	v1 := app.Group("/api/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {
		return h.LoginHandler(c, ctx, ls)
	})

	v1.Get("/user/:email", func(c *fiber.Ctx) error {
		return h.GetUsersHandler(c, ctx, ls)
	})

	v1.Post("/register", func(c *fiber.Ctx) error {
		return h.RegisterHandler(c, ctx, ls)
	})

	v1.Delete("/delete", func(c *fiber.Ctx) error {
		return h.DeleteHandler(c, ctx, ls)
	})

	v1.Patch("/update", func(c *fiber.Ctx) error {
		return h.UpdateHandler(c, ctx, ls)
	})
}
