package dto

import "time"

type CreateUserRequest struct {
	Id            string    `json:"id,omitempty"`
	Email         string    `json:"email" validate:"required,email"`
	Phone         string    `json:"phone" validate:"required,e164"`
	Password      string    `json:"password" validate:"required,min=6"`
	EmailVerified bool      `json:"email_verified,omitempty"`
	IsActive     bool      `json:"is_active,omitempty"`
	IsBlocked    bool      `json:"is_blocked,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	Name          string    `json:"name,omitempty"`
}
