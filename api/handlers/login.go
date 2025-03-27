package handlers

import (
	"context"
	"encoding/json"
	"go-lang/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"

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
func LoginHandler(c *fiber.Ctx, ctx context.Context, r repositories.RedisClient) error {
	var req LoginRequest

	// Validação do request
	err := IsValidLoginRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	// Lógica de login
	data, err := r.GetUser(ctx, req.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}

	if data.Password == req.Password {
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
func RegisterHandler(c *fiber.Ctx, ctx context.Context, r repositories.RedisClient) error {
	var req repositories.User

	err := IsValidCreateUserRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	data, err := r.GetUser(ctx, req.Email)
	if err != nil && err != redis.Nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if data != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Usuário já existe")
	}

	value, err := json.Marshal(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao serializar dados: "+err.Error())
	}

	err = r.Set(ctx, req.Email, value)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao inserir dados: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário registrado com sucesso"})
}
