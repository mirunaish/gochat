package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// connect to database
func Connect() error {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))	
	var err error

	// connect to postgres
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	// auto migrate all models
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("database: failed to migrate models")
		return err
	}

	return nil
}

