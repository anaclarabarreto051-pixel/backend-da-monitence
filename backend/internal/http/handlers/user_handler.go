package handlers

import (
	"context"
	"net/http"

	"myapp/internal/domain"
	"myapp/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type UserCreator interface {
	Create(ctx context.Context, user *domain.User) error
}

type Handler struct {
	UserRepo UserCreator
}

func NewHandler(userRepo UserCreator) *Handler {
	return &Handler{UserRepo: userRepo}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Email:        req.Email,
		Phone:        req.Phone,
		Name:         req.Name,
		PasswordHash: req.Password,
	}

	if err := h.UserRepo.Create(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": user.ID})
}