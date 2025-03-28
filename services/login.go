package services

import (
	"context"
	"encoding/json"
	"go-lang/repositories"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type LoginService struct {
	redis *repositories.RedisClient
}

func NewLoginService(r *repositories.RedisClient) *LoginService {
	return &LoginService{
		redis: r,
	}
}

func (l *LoginService) LoginUser(ctx context.Context, email, password string) (bool, error) {
	data, err := l.redis.GetUser(ctx, email)
	if err != nil {
		return false, fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}

	if data == nil {
		return false, fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
	}

	if data.Password == password {
		return true, nil
	}

	return false, fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
}

func (l *LoginService) CreateUser(ctx context.Context, req *repositories.User) error {
	data, err := l.redis.GetUser(ctx, req.Email)
	if err != nil && err != redis.Nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if data != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Usuário já existe")
	}

	value, err := json.Marshal(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao serializar dados: "+err.Error())
	}

	err = l.redis.Set(ctx, req.Email, value)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao inserir dados: "+err.Error())
	}

	return nil
}

func (l *LoginService) DeleteUser(ctx context.Context, req *repositories.User) error {
	data, err := l.redis.GetUser(ctx, req.Email)
	if err != nil && err != redis.Nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if data == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Usuário não encontrado")
	}

	if data.Password != req.Password {
		return fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
	}

	err = l.redis.Del(ctx, req.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao deletar dados: "+err.Error())
	}

	return nil
}

func (l *LoginService) UpdateUser(ctx context.Context, req *repositories.User) error {
	data, err := l.redis.GetUser(ctx, req.Email)
	if err != nil && err != redis.Nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if data == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Usuário não encontrado")
	}

	if data.Password != req.Password {
		return fiber.NewError(fiber.StatusUnauthorized, "Credenciais inválidas")
	}

	value, err := json.Marshal(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao serializar dados: "+err.Error())
	}

	err = l.redis.Set(ctx, req.Email, value)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Erro ao inserir dados: "+err.Error())
	}

	return nil
}

func (l *LoginService) GetUser(ctx context.Context, email string) (*repositories.User, error) {
	data, err := l.redis.GetUser(ctx, email)
	if err != nil && err != redis.Nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar dados: "+err.Error())
	}
	if data == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Usuário não encontrado")
	}

	return data, nil
}
