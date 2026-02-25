package http

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/http/handlers"
)
func RegisterRoutes(r *gin.Engine, h *handlers.Handler	) {
	api := r.Group("/api")
	{
		api.POST("/users", h.CreateUser)
	}
}
