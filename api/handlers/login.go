package handlers

import (
	"context"
	mongo "go-lang/repositories"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"

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
func LoginHandler(c *fiber.Ctx, ctx *context.Context) error {
	var req LoginRequest

	// Validação do request
	err := IsValidLoginRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	// Lógica de login
	filter := bson.M{"email": req.Email, "password": req.Password}
	data, err := mongo.ReadItems(ctx, filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}

	for _, item := range data {
		if item.Email == req.Email && item.Password == req.Password {
			return c.JSON(fiber.Map{"message": "Login efetuado com sucesso"})
		}
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
}

// RegisterHandler é o handler para a rota POST /api/v1/register
// @Summary Register
// @Description Registra um novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body mongo.User true "Corpo da requisição"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/register [post]
func RegisterHandler(c *fiber.Ctx, ctx *context.Context) error {
	// Extrai o corpo da requisição
	var req mongo.User
	err := IsValidCreateUserRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	// Lógica de registro

	// Checa se o usuário já existe
	filter := bson.M{"email": req.Email}
	data, err := mongo.ReadItems(ctx, filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if len(data) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Usuário já existe")
	}

	// Insere o usuário
	err = mongo.InsertItem(ctx, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao inserir dados: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário registrado com sucesso"})
}
