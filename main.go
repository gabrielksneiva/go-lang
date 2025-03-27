package main

import (
	"context"
	"go-lang/api"
	_ "go-lang/docs"
	"go-lang/repositories"
	"log"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Inicializa o app usando a função New do pacote api

	ctx := context.Background()

	cache, err := repositories.NewRedisClient()
	if err != nil {
		log.Fatalf("Erro ao criar o cliente Redis: %v", err)
	}

	app := api.NewApp(ctx, *cache)

	// Inicia o container do MongoDB e obtém a URI
	// mongoURI := repositories.StartMongoContainer()

	// Conecta ao MongoDB
	// repositories.ConnectToMongo(mongoURI)

	// Configura o Swagger
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// Inicia o servidor na porta 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
