package main

import (
	"log"
	"os"

	drinkHttp "goexample/internal/drink/delivery/http"
	drinkPostgres "goexample/internal/drink/repository/postgres"
	drinkUsecase "goexample/internal/drink/usecase"
	userHttp "goexample/internal/user/delivery/http"
	userPostgres "goexample/internal/user/repository/postgres"
	userUsecase "goexample/internal/user/usecase"
	"goexample/pkg/db"
	"goexample/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	database := db.NewPostgresDB()

	userRepo := userPostgres.NewPostgresUserRepository(database)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userHttp.NewUserHandler(userUsecase)

	drinkRepo := drinkPostgres.NewPostgresDrinkRepository(database)

	accessKey := os.Getenv("R2_ACCESS_KEY")
	secretKey := os.Getenv("R2_SECRET_KEY")
	accountID := os.Getenv("R2_ACCOUNT_ID")
	bucketName := os.Getenv("R2_BUCKET_NAME")

	if accessKey == "" || secretKey == "" || accountID == "" || bucketName == "" {
		log.Fatal("As credenciais do Cloudflare R2 não estão definidas nas variáveis de ambiente")
	}

	r2Storage, err := storage.NewR2Storage(accessKey, secretKey, accountID, bucketName)
	if err != nil {
		log.Fatalf("Erro ao inicializar o R2Storage: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	drinkUsecase := drinkUsecase.NewDrinkUsecase(drinkRepo, r2Storage)
	drinkHandler := drinkHttp.NewDrinkHandler(drinkUsecase)

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bem-vindo à API",
		})
	})

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.ListUsers)

	router.POST("/drinks", drinkHandler.CreateDrink)
	router.GET("/drinks", drinkHandler.ListDrinks)
	router.DELETE("/drinks", drinkHandler.DeleteDrink)

	log.Printf("Servidor rodando na porta %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
