package services

import (
	"log"

	"github.com/google/uuid"

	"github.com/mirunaish/gochat/server/internal/database"
	"github.com/mirunaish/gochat/server/internal/models"
	"github.com/mirunaish/gochat/server/internal/socket"
	"github.com/mirunaish/gochat/server/internal/utils"
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

// get active users. exclude self.
func GetActiveUsers(userId string) ([]*models.User, error) {
	subscribers := socket.GetAllSubscribers()

	activeUsers := []*models.User{}
	// loop over subscribers and append user information
	for _, sub := range subscribers {
		// skip myself
		if userId == sub.UserId {
			continue
		}

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
