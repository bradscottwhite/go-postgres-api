package database

import (
	"log"
	"fmt"
	"os"
	"github.com/joho/godotenv"

	"github.com/bradscottwhite/go-postgres-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

// connectDb
func ConnectDb() {
	var (
		host	 = getEnvVar("DB_HOST")
		port	 = getEnvVar("DB_PORT")
		user	 = getEnvVar("DB_USER")
		dbname	 = getEnvVar("DB_NAME")
		password = getEnvVar("DB_PASSWORD")
	)
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}
}
