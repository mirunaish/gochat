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

type DatabaseInfo struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

// connect to database
func Connect() error {
	var db_info DatabaseInfo

	// if these env variables are set, we are in production environment. use rds
	if os.Getenv("RDS_HOSTNAME") != "" {
		db_info = DatabaseInfo{
			Host: os.Getenv("RDS_HOSTNAME"),
			Port: os.Getenv("RDS_PORT"),
			User: os.Getenv("RDS_USERNAME"),
			Pass: os.Getenv("RDS_PASSWORD"),
			Name: os.Getenv("RDS_DB_NAME"),
		}
	} else {
		// otherwise use values given in .env
		db_info = DatabaseInfo{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		}
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		db_info.Host, db_info.User, db_info.Pass, db_info.Name, db_info.Port)
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
	}

	return nil
}
