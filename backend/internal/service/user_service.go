package service

import (
	"context"
	"errors"
	"strings"

	"myapp/internal/http/dto"
	"myapp/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}