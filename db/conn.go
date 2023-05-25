package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/joho/godotenv"
)

func DbInstance() *database.Queries {
	//database connection
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("No database connection found")
	}

	fmt.Println(db_url)
	db, dbErr := sql.Open("postgres", db_url)
	if dbErr != nil {
		log.Fatal("Failed to connect to database")
	}

	dbQueries := database.New(db)
	return dbQueries
}

var DbClient *database.Queries = DbInstance()
