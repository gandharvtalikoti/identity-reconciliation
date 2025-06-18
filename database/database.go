package database

import (
	"bitespeed-identity/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	DB = db
	fmt.Println("✅ Database connection successful")
}
