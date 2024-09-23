package utils

import (
	"encoding/base64"
	"log"
	"os"
	"time"

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
		log.Print("auth service: failed to decode jwt key (check your environment variable)")
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

	// https://datatracker.ietf.org/doc/html/rfc7519
	claims := jwt.MapClaims{"iss": os.Getenv("ISSUER"), "alg": "HS256", "sub": userId, "iat": time.Now().Unix()}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		log.Print("auth service: failed to sign jwt")
		return "", err
	}

	return token, nil
}

// https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
// return the user id and true if the token is valid
func ParseAndVerifyJwt(tokenSigned string) (string, bool) {
	key, err := getJwtKey()
	if err != nil {
		log.Print("auth: failed to get jwt key. check your environment variable")
		return "", false
	}

	claims := &jwt.RegisteredClaims{}
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}

	_, err = jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}), jwt.WithIssuer(os.Getenv("ISSUER"))).ParseWithClaims(tokenSigned, claims, keyfunc)
	if err != nil {
		log.Print("auth: failed to parse jwt")
		return "", false
	}

	userId, err := claims.GetSubject()

	// if no error, token was valid
	return userId, err == nil
}
