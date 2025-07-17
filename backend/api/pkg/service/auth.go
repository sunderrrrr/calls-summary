package service

import "api/pkg/repository"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) signUp(email, password string) (int, error) {
	//mock
	return 1, nil
}

func (s *AuthService) signIn(email, password string) (string, error) {
	//mock
	return "token", nil
}
