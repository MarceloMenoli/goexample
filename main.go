package main

import (
	"fmt"
	"log"

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

	// Exemplo de inserção (descomente para testar)
	//_, err = db.Exec("INSERT INTO usuarios (nome, email) VALUES ('Fulano', 'fulano@email.com')")
	//if err != nil {
	//	log.Fatalf("Erro ao inserir dados: %v", err)
	//}
	//fmt.Println("Dados inseridos com sucesso!")
}
