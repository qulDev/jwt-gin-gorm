package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qulDev/jwt-gin-gorm/internal/handler"
)

type Handlers struct {
	User *handler.UserHandler
}

func SetupRoutes(h *Handlers) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	AuthRoutes(v1, h.User)

	return r
}
