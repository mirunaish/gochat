package services

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
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

	// hash password
	password, err := hashPassword(password)
	if err != nil {
		log.Fatalf("user service: failed to hash password: %s", err.Error())
		return nil, err
	}

	// create user
	newUser := models.User{ ID: uuid.New().String(), Email: email, Username: username, Password: password }
	err = database.CreateUser(&newUser)
	if err != nil {
		// TODO better error messages here
		log.Fatalf("user service: failed to create user: %s", err.Error())
		return nil, err
	}

	return &newUser, nil
}

func LogIn(email, password string) (string, error) {
	// find user with email
	user, err := database.GetUserByEmail(email) // TODO
	if err != nil {
		return "", &utils.RouterError{Code: http.StatusNotFound, Message: "no account found with that email" }
	}

	// get password hash from db
	hash := user.Password

	// check that hash matches given password
	if !checkPassword(hash, password) {
		return "", &utils.RouterError{Code: http.StatusBadRequest, Message: "password is incorrect"}
	}

	return "", nil
}