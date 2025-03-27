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
	ctx := context.Background()

	cache, err := repositories.NewRedisClient()
	if err != nil {
		log.Fatalf("Erro ao criar o cliente Redis: %v", err)
	}

	app := api.NewApp(ctx, *cache)

	// mongoURI := repositories.StartMongoContainer()

	// repositories.ConnectToMongo(mongoURI)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
