package service

import "api/pkg/repository"

type Service struct {
	Auth
}

type Auth interface {
	signUp(email, password string) (int, error)
	signIn(email, password string) (string, error)
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository.Auth),
	}
}
