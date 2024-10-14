// pkg/db/postgres.go
package db

import (
	"log"
	"os"

	"goexample/internal/drink"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_PUBLIC_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	if err := db.AutoMigrate(&drink.Drink{}); err != nil {
		log.Fatal("Falha ao migrar a tabela de drinks:", err)
	}
	return db
}
