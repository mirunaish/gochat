package services

import (
	"log"
	"net/http"

	"github.com/mirunaish/gochat/server/internal/database"
	"github.com/mirunaish/gochat/server/internal/utils"
)

// find user and return a jwt
func LogIn(email, password string) (string, error) {
	// find user with email
	user, err := database.GetUserByEmail(email)
	if err != nil {
		return "", &utils.RouterError{Code: http.StatusNotFound, Message: "no account found with that email"}
	}

	// get password hash from db
	hash := user.Password

	// check that hash matches given password
	if !utils.CheckPassword(hash, password) {
		return "", &utils.RouterError{Code: http.StatusBadRequest, Message: "password is incorrect"}
	}

	// create jwt
	jwt, err := utils.CreateJwt(user.ID, user.Email)
	if err != nil {
		log.Printf("auth service: failed to create jwt: %s", err.Error())
		return "", &utils.RouterError{Code: http.StatusBadRequest, Message: "failed to log in"}
	}

	return jwt, nil
}

// create user and return a jwt
func SignUp(email, username, password string) (string, error) {
	// create user
	newUser, err := CreateUser(email, username, password)
	if err != nil {
		log.Printf("routes: failed to create new user: %s", err.Error())
		return "", err
	}

	// attempt to log user in
	jwt, err := utils.CreateJwt(newUser.ID, newUser.Email)
	if err != nil {
		log.Printf("routes: failed to log in after signup: %s", err.Error())
		// create new error with better message for user
		err = &utils.RouterError{Code: http.StatusCreated, Message: "user was created, but failed to log in. please try to log in"}
		return "", err
	}

	return jwt, nil
}
