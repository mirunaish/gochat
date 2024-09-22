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
