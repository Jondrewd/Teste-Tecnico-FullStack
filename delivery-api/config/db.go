package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB inicializa e retorna uma conexão com o banco de dados MySQL.
// A função configura a string de conexão (DSN) e tenta estabelecer a conexão.
// Se a conexão falhar, o programa é encerrado com um erro fatal.
func InitDB() *gorm.DB {
	// Define a string de conexão (DSN) para o MySQL.
	// Formato: "usuário:senha@tcp(endereço:porta)/nome_do_banco?parâmetros_adicionais"
	dsn := "root:1234@tcp(localhost:3306)/desafio?charset=utf8mb4&parseTime=True&loc=Local"

	// Tenta estabelecer a conexão com o banco de dados usando o driver MySQL.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Se a conexão falhar, loga um erro fatal e encerra o programa.
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Retorna a instância do GORM (*gorm.DB) que representa a conexão com o banco de dados.
	return db
}