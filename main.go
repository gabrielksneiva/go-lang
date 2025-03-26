package main

import (
	"go-lang/api"
	_ "go-lang/docs"
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Inicializa o app usando a função New do pacote api
	app := fiber.New()
	api.Routes(app)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// Inicia o servidor na porta 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
