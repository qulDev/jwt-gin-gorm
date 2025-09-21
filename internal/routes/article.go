package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qulDev/jwt-gin-gorm/internal/handler"
)

func ArticleRoute(r *gin.RouterGroup, h *handler.ArticleHandler) {

	articles := r.Group("/articles")
	{
		articles.GET("/", h.GetArticles)
	}
}
