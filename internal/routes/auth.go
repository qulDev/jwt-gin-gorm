package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qulDev/jwt-gin-gorm/internal/handler"
)

func AuthRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.RegisterUser)
	}
}
