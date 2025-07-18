package service

import (
	"api/models"
	"api/pkg/repository"
)

type Service struct {
	Auth
}

type Auth interface {
	SignUp(input models.SignUpInput) (int, error)           // register
	GenerateToken(input models.SignInInput) (string, error) // login
	ParseToken(token string) (models.User, error)           // middleware
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository.Auth),
	}
}
