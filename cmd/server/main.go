package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

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

	// Inicializa o servidor Gin
	router := gin.Default()

	// Definir a porta (pegando da variável de ambiente PORT ou definindo padrão 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Rotas para os usuários
	router.POST("/users", userHandler.CreateUser) // Rota para criar usuário
	router.GET("/users", userHandler.ListUsers)   // Rota para listar usuários

	// Iniciar o servidor na porta definida
	log.Printf("Servidor rodando na porta %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
