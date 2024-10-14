package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// URL de conexão com o PostgreSQL (Railway)
	dsn := "postgresql://postgres:BZasVXAJIVaLTVMTGnwVZQppJgexVnUQ@autorack.proxy.rlwy.net:58916/railway"

	// Conectar ao banco de dados usando GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar o modelo para criar a tabela "users"
	db.AutoMigrate(&User{})

	fmt.Println("Conexão ao PostgreSQL com GORM realizada com sucesso!")

	// Definir a porta (pegando da variável de ambiente PORT ou definindo padrão 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Rota simples para verificar se o servidor está funcionando
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Conexão ao PostgreSQL com GORM realizada com sucesso!")
	})

	// Iniciar o servidor HTTP na porta definida
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
