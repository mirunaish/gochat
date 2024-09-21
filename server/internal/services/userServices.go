package services

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// helper function to hash a password before saving user to db
func hashPassword(password string) (string, error) {
	// convert password to bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err  // convert bytes back to string
}

// helper function to check a password against a password hash
func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// create a user
func CreateUser(email, username, password string) (*models.User, error) {
	// check inputs TODO

	password, err := hashPassword(password)
	if err != nil {return nil, err}

	// TODO guid
	newUser := models.User{ ID: "0", Email: email, Username: username, Password: password }
	
	// if email already exists etc, can "throw" a routerError. example
	return &newUser, &utils.RouterError{ Code: http.StatusConflict, Message: "email already exists" }

	// TODO create in database
	return &newUser, nil
}

func LogIn(email, password string) (string, error) {
	// find user with email
	user := models.User{} // TODO

	// get password hash from db
	hash := user.Password

	// check that hash matches given password
	if !checkPassword(hash, password) {
		return "", &utils.RouterError{Code: http.StatusBadRequest, Message: "Password is incorrect"}
	}

	return "", nil
}