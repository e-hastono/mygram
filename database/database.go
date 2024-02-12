package database

import (
	"fmt"
	"log"
	"os"

	"github.com/e-hastono/mygram/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("[gorm database] Successfully connected to Postgres database %s on %s:%s\n\n", dbName, dbHost, dbPort)
	}

	db.Debug().AutoMigrate(&entities.User{}, &entities.SocialMedia{}, &entities.Photo{}, &entities.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
