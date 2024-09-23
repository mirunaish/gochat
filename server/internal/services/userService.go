package services

import (
	"log"

	"github.com/google/uuid"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/socket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// create a user
func CreateUser(email, username, password string) (*models.User, error) {
	// check inputs TODO

	// hash password
	password, err := utils.HashPassword(password)
	if err != nil {
		log.Printf("user service: failed to hash password: %s", err.Error())
		return nil, err
	}

	// create user
	// https://pkg.go.dev/github.com/google/uuid
	newUser := models.User{ID: uuid.New().String(), Email: email, Username: username, Password: password}
	err = database.CreateUser(&newUser)
	if err != nil {
		// TODO better error messages here
		log.Printf("user service: failed to create user: %s", err.Error())
		return nil, err
	}

	return &newUser, nil
}

func GetActiveUsers() ([]*models.User, error) {
	subscribers := socket.GetAllSubscribers()

	activeUsers := []*models.User{}
	// loop over subscribers and append user information
	for _, sub := range subscribers {
		// get user with this user id
		user, err := database.GetUser(sub.UserId)
		if err != nil {
			log.Printf("user service: failed to get users: %s", err.Error())
			return nil, err
		}

		activeUsers = append(activeUsers, user)
	}

	return activeUsers, nil
}
