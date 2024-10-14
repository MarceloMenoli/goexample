// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"os"

	userHttp "goexample/internal/user/delivery/http"
	"goexample/internal/user/repository/postgres"
	"goexample/internal/user/usecase"
	"goexample/pkg/db"
)

func main() {
	// Inicializa o banco de dados
	database := db.NewPostgresDB()

	// Inicializa o repositório, caso de uso e o handler
	userRepo := postgres.NewPostgresUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := userHttp.NewUserHandler(userUsecase)

	// Definir a porta (pegando da variável de ambiente PORT ou definindo padrão 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configurar rotas
	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/list", userHandler.ListUsers)

	// Iniciar o servidor HTTP na porta definida
	log.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
