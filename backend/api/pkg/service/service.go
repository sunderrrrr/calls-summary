package service

import (
	"api/models"
	"api/pkg/repository"
	"io"
	"time"
)

type Service struct {
	Auth
	User
	Report
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

type User interface {
	ResetPassword(request models.UserReset) error
	ForgotPassword(request models.ResetRequest) error
}

type Report interface {
	ReportCall(file io.Reader, filename string) (string, error)
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth:   NewAuthService(repository.Auth),
		Report: NewReportService(),
		User:   NewUserService(repository.User),
	}
}
