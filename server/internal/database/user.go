package database

import "github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"

// returns only id and username for each user
func GetAllUsers(id string) (*[]models.User, error) {
	var users []models.User
	result := db.Select("id", "username").Find(&users)
	return &users, result.Error
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	result := db.First(&user, "id = ?", id)
	return &user, result.Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := db.First(&user, "email = ?", email)
	return &user, result.Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := db.First(&user, "username = ?", username)
	return &user, result.Error
}

func CreateUser(user *models.User) error {
	result := db.Create(user)
	return result.Error
}
