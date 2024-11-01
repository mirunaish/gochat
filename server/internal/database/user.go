package database

import (
	"errors"

	"github.com/mirunaish/gochat/server/internal/models"
)

// fake database: map in server memory.
// violates rest principles (specifically statelessness)
var users = make(map[string]models.User)

// returns only id and username for each user
// exclude self
// unused
// func GetAllUsers(userId string) (*[]models.User, error) {
// 	var users []models.User
// 	result := db.Select("id", "username").Not("id = ?", userId).Find(&users)
// 	return &users, result.Error
// }

func GetUser(id string) (*models.User, error) {
	var user models.User
	// result := db.First(&user, "id = ?", id)
	user = users[id]

	// if user not found, return error
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	// return &user, result.Error
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	// result := db.First(&user, "email = ?", email)
	// find user with email in map
	for _, u := range users {
		if u.Email == email {
			user = u
			break
		}
	}

	// if user not found, return error
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	// return &user, result.Error
	return &user, nil
}

// unused
// func GetUserByUsername(username string) (*models.User, error) {
// 	var user models.User
// 	result := db.First(&user, "username = ?", username)
// 	return &user, result.Error
// }

func CreateUser(user *models.User) error {
	// result := db.Create(user)
	users[user.ID] = *user
	// return result.Error
	return nil
}
