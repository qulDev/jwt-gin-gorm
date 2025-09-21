package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(h *Handlers) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	AuthRoutes(v1, h.User)
	ArticleRoute(v1, h.Article)

	return r
}
