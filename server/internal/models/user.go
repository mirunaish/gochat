package models

import "time"

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"` // encoded
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"` // unencoded
}