package main

import (
	"context"
	"go-lang/api"
	_ "go-lang/docs"
	mongo "go-lang/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Inicializa o app usando a função New do pacote api
	app := fiber.New()
	ctx := context.Background()

	// Inicia o container do MongoDB e obtém a URI
	mongoURI := mongo.StartMongoContainer()

	// Conecta ao MongoDB
	mongo.ConnectToMongo(mongoURI)

	// Configura as rotas
	api.Routes(&ctx, app)

	// Configura o Swagger
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// Inicia o servidor na porta 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
