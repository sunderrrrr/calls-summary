package service

import (
	"api/models"
	"api/pkg/repository"
	"time"
)

type Service struct {
	Auth
}

const (
	salt          = "hifu&hfI&fG&Igaw"
	secretKey     = "2c3982433nc89m43v3n89323492u49"
	tokenTTL      = 12 * time.Hour
	resetTokenTTL = time.Minute * 10
)

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
