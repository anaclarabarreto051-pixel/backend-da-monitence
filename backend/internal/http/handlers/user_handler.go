package handlers

import (
	"myapp/internal/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)		

type Handler struct {
	// You can add dependencies here, such as a service layer or database connection
}
func (h *Handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}