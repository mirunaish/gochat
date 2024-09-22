package utils

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// helper function to hash a password before saving user to db
func HashPassword(password string) (string, error) {
	// convert password to bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err // convert bytes back to string
}

// helper function to check a password against a password hash
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// return the jwt key environment variable as bytes
func getJwtKey() ([]byte, error) {
	key := os.Getenv("JWT_KEY")
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Fatalf("auth service: failed to decode jwt key (check your environment variable)")
		return nil, err
	}

	return keyBytes, nil
}

// create a jwt with user data encoded in claims
func CreateJwt(userId, email string) (string, error) {
	key, err := getJwtKey()
	if err != nil {
		return "", err
	}

	// TODO add more claims?
	claims := jwt.MapClaims{"alg": "HS256"}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		log.Fatalf("auth service: failed to sign jwt")
		return "", err
	}

	return token, nil
}

// https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
func ParseAndVerifyJwt(token string) bool {
	key, err := getJwtKey()
	if err != nil {
		return false
	}

	// TODO

	claims := &jwt.RegisteredClaims{}

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}) // .WithValidMethods([]string{"HS256"})
	return true
}
