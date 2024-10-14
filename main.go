package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

var db *gorm.DB

func main() {
	// URL de conexão com o PostgreSQL (Railway)
	dsn := "postgresql://postgres:BZasVXAJIVaLTVMTGnwVZQppJgexVnUQ@autorack.proxy.rlwy.net:58916/railway"

	// Conectar ao banco de dados usando GORM
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	// Rota principal para teste de conexão
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Conexão ao PostgreSQL com GORM realizada com sucesso!")
	})

	// Rota para criar novo usuário
	http.HandleFunc("/users", createUser)

	// Iniciar o servidor HTTP na porta definida
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Função para lidar com a criação de usuário
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar o corpo da requisição
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Validar se o campo "Name" não está vazio
	if user.Name == "" {
		http.Error(w, "O campo 'name' é obrigatório", http.StatusBadRequest)
		return
	}

	// Criar o novo usuário no banco de dados
	if result := db.Create(&user); result.Error != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	// Retornar a resposta com o ID do usuário criado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
