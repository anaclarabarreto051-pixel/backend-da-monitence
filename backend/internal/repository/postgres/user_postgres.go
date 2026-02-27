package postgres

import (
	"database/sql"
		"context"
	"myapp/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (email, phone, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`
  return r.db.QueryRowContext(ctx, query, user.Email, user.Phone, user.PasswordHash).Scan(&user.ID)
}

