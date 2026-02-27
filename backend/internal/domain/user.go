package domain

import "time"

type User struct {
	ID        int64    `json:"id"`
	Email     string    `json:"email"`
	Phone    string    `json:"phone"`
	PasswordHash  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}