package main

import (
	"api-go-lang/api"
	"log"
)

func main() {
	// Inicializa o app usando a função New do pacote api
	app := api.New()

	// Inicia o servidor na porta 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
