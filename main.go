package main

import (
	"context"
	"go-lang/api"
	_ "go-lang/docs"
	"log"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	ctx := context.Background()

	app := api.NewApp(ctx)

	// mongoURI := repositories.StartMongoContainer()

	// repositories.ConnectToMongo(mongoURI)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
