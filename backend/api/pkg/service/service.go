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
	Analysis
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

type Analysis interface {
	AnalyzeCall(userId int, file io.Reader, filename string) (string, error)
	SendMessageToChat(analysisId string, userId int, message models.ChatMessage) error
	GetChatHistory(analysisId string, userId int) ([]models.ChatMessage, error)
}

type Chat interface {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth:     NewAuthService(repository.Auth),
		User:     NewUserService(repository.User),
		Analysis: NewAnalysisService(repository.Analysis),
	}
}
