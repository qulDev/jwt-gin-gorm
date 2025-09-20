package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qulDev/jwt-gin-gorm/internal/dto"
	"github.com/qulDev/jwt-gin-gorm/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"messege": "success register user",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
