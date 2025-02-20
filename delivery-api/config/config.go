package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente de um arquivo .env.
// Se o arquivo .env não for encontrado, o sistema usará as variáveis de ambiente do sistema operacional.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

// GetDatabaseConnection estabelece uma conexão com o banco de dados com base no tipo de banco de dados configurado.
// Ele suporta SQLite, PostgreSQL e MySQL.
// Retorna uma instância do GORM (*gorm.DB) ou um erro, caso a conexão falhe.
func GetDatabaseConnection() (*gorm.DB, error) {
	// Carrega as variáveis de ambiente do arquivo .env ou do sistema.
	LoadEnv()

	// Define o tipo de banco de dados a ser utilizado.
	// Aqui, foi definido diretamente como "mysql" para facilitar a execução pelos recrutadores.
	// Em um ambiente real, você pode usar: dbType := os.Getenv("DATABASE_TYPE")
	dbType := "mysql"

	var db *gorm.DB
	var err error

	// Escolhe o driver de banco de dados com base no tipo configurado.
	switch dbType {
	case "sqlite":
		// Configuração para SQLite (banco de dados embutido).
		// O banco de dados será criado no arquivo "gorm.db".
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	case "postgres":
		// Configuração para PostgreSQL.
		// Monta a string de conexão (DSN) usando as variáveis de ambiente.
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),     // Endereço do servidor PostgreSQL
			os.Getenv("DB_USER"),     // Usuário do banco de dados
			os.Getenv("DB_PASSWORD"), // Senha do banco de dados
			os.Getenv("DB_NAME"),     // Nome do banco de dados
			os.Getenv("DB_PORT"),     // Porta do servidor PostgreSQL
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		// Configuração para MySQL.
		// Monta a string de conexão (DSN) usando as variáveis de ambiente.
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),     // Usuário do banco de dados
			os.Getenv("DB_PASSWORD"), // Senha do banco de dados
			os.Getenv("DB_HOST"),     // Endereço do servidor MySQL
			os.Getenv("DB_PORT"),     // Porta do servidor MySQL
			os.Getenv("DB_NAME"),     // Nome do banco de dados
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		// Retorna um erro se o tipo de banco de dados não for suportado.
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	// Verifica se houve erro ao conectar ao banco de dados.
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Loga uma mensagem de sucesso ao estabelecer a conexão.
	log.Println("Database connection established successfully")
	return db, nil
}