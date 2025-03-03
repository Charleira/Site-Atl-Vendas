package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // driver MySQL
	"github.com/joho/godotenv"
)

// DB é a variável global que irá armazenar a conexão com o banco de dados
var DB *sql.DB

// ConnectToDatabase é responsável por inicializar a conexão com o banco de dados
func ConnectToDatabase() {
	// Carregar variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Construir a string de conexão com os dados do .env
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	// Conectar ao banco de dados com *sql.DB
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Testar a conexão
	if err := DB.Ping(); err != nil {
		log.Fatalf("Erro ao testar a conexão com o banco de dados: %v", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
}
