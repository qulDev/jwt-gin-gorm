package service

import (
	"errors"

	"github.com/qulDev/jwt-gin-gorm/internal/models"
	"github.com/qulDev/jwt-gin-gorm/internal/repository"
	"github.com/qulDev/jwt-gin-gorm/pkg/hash"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) Register(email, password string) (*models.User, error) {
	//check if user already exists

	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// if err not from record not found, return the error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//	hash password
	hashPassword, err := hash.HashPassword(password)

	user := &models.User{
		Email:        email,
		PasswordHash: hashPassword,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
