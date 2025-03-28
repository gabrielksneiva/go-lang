package handlers

import (
	"context"
	"go-lang/repositories"
	s "go-lang/services"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// LoginHandler é o handler para a rota POST /api/v1/login
// @Summary Login
// @Description Realiza o login de um usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Corpo da requisição"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/login [post]
func LoginHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req LoginRequest

	err := IsValidLoginRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	ok, err := ls.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao realizar login: "+err.Error())
	}

	if ok {
		return c.JSON(fiber.Map{"message": "Login realizado com sucesso"})
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
}

// RegisterHandler é o handler para a rota POST /api/v1/register
// @Summary Register
// @Description Registra um novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body repositories.User true "Corpo da requisição"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/register [post]
func RegisterHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req repositories.User
	err := IsValidCreateUserRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	err = ls.CreateUser(ctx, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao registrar usuário: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário registrado com sucesso"})
}

func DeleteHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req DeleteRequest
	var user repositories.User
	err := IsValidDeleteRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	err = ls.DeleteUser(ctx, &user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao deletar usuário: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário deletado com sucesso"})
}
