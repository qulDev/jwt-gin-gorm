package handler

import "github.com/gin-gonic/gin"

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	c.JSON(200, gin.H{"message": "all articles"})
}
