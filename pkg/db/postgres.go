// pkg/db/postgres.go
package db

import (
	"fmt"
	"log"
	"os"

	"goexample/internal/drink"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar o arquivo .env")
		}
	}

	dsn := os.Getenv("DATABASE_PUBLIC_URL")
	if dsn == "" {
		log.Fatal("DATABASE_PUBLIC_URL não está definida")
	}
	fmt.Println("Conectando ao banco de dados com DSN:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	if err := db.AutoMigrate(&drink.Drink{}); err != nil {
		log.Fatal("Falha ao migrar a tabela de drinks:", err)
	}
	return db
}
