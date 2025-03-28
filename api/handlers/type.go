package handlers

import (
	mongo "go-lang/repositories"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DeleteRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type GetRequest struct {
	Email string `json:"email" validate:"required,email"`
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

func IsValidDeleteRequest(c *fiber.Ctx, req *DeleteRequest) error {
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}

func IsValidUpdateRequest(c *fiber.Ctx, req *UpdateRequest) error {
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}

func IsValidGetUsersRequest(c *fiber.Ctx, req *GetRequest) error {
	encodedEmail := c.Params("email")
	if encodedEmail == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email não informado")
	}

	decodedEmail, err := url.QueryUnescape(encodedEmail)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro ao decodificar o email")
	}

	req.Email = decodedEmail
	return nil
}
