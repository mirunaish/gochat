package services

import (
	"log"

	"github.com/google/uuid"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// create a user
func CreateUser(email, username, password string) (*models.User, error) {
	// check inputs TODO

	// hash password
	password, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("user service: failed to hash password: %s", err.Error())
		return nil, err
	}

	// create user
	// https://pkg.go.dev/github.com/google/uuid
	newUser := models.User{ID: uuid.New().String(), Email: email, Username: username, Password: password}
	err = database.CreateUser(&newUser)
	if err != nil {
		// TODO better error messages here
		log.Fatalf("user service: failed to create user: %s", err.Error())
		return nil, err
	}

	return &newUser, nil
}
