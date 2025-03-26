package handlers

import (
	mongo "go-lang/repositories"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func IsValidLoginRequest(c *fiber.Ctx, req *LoginRequest) error {
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}

func IsValidCreateUserRequest(c *fiber.Ctx, req *mongo.User) error {
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}
