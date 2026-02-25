package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Phone    string    `json:"phone"`
	Password_hash  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}