package models

type SignupRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"` // unencoded
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"` // unencoded
}