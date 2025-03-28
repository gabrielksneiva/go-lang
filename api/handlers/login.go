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

// DeleteHandler é o handler para a rota DELETE /api/v1/delete
// @Summary Delete
// @Description Deleta um usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body DeleteRequest true "Corpo da requisição"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/delete [delete]
func DeleteHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req DeleteRequest
	err := IsValidDeleteRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	userInfosToCheck := repositories.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err = ls.DeleteUser(ctx, &userInfosToCheck)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao deletar usuário: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário deletado com sucesso"})
}

// UpdateHandler é o handler para a rota PATCH /api/v1/update
// @Summary Update
// @Description Atualiza um usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body UpdateRequest true "Corpo da requisição"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/update [patch]
func UpdateHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req UpdateRequest
	err := IsValidUpdateRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	userInfosToUpdate := repositories.User{
		Email:    req.Email,
		Password: req.Password,
		ID:       req.ID,
		Name:     req.Name,
	}
	err = ls.UpdateUser(ctx, &userInfosToUpdate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao atualizar usuário: "+err.Error())
	}

	return c.JSON(fiber.Map{"message": "Usuário atualizado com sucesso"})
}

// GetUsersHandler é o handler para a rota GET /api/v1/user/:email
// @Summary Get Users
// @Description Busca um usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param email path string true "Email do usuário"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/user/{email} [get]
func GetUsersHandler(c *fiber.Ctx, ctx context.Context, ls *s.LoginService) error {
	var req GetRequest
	err := IsValidGetUsersRequest(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Erro de validação: "+err.Error())
	}

	users, err := ls.GetUser(ctx, req.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar usuários: "+err.Error())
	}

	return c.JSON(fiber.Map{"users": users})
}
