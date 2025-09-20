package service

import (
	"fmt"

	"github.com/qulDev/jwt-gin-gorm/internal/models"
	"github.com/qulDev/jwt-gin-gorm/internal/repository"
	"github.com/qulDev/jwt-gin-gorm/pkg/hash"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) Register(email, password string) (*models.User, error) {
	//	hash password
	hashedPassword, errHash := hash.HashPassword(password)
	if errHash != nil {
		return nil, fmt.Errorf("failed to hash password: %w", errHash)
	}

	user := &models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}
