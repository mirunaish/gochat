package services

import (
	"net/http"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// create a user
func CreateUser(email string, username string, password string) (models.User, error) {
	// TODO guid, encrypt password, etc
	newUser := models.User{ ID: "0", Email: email, Username: username, Password: password }

	// if email already exists etc, can "throw" a routerError. example
	return newUser, &utils.RouterError{ Code: http.StatusConflict, Message: "email already exists" }

	// TODO create in database
	return newUser, nil
}