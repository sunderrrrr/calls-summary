package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
	User
	Analysis
}
type Auth interface {
	SignUp(email, name, password string) (int, error)
	GetUser(email, password string) (models.User, error)
}
type User interface {
	ResetPassword(email, newPassword string) error
}

type Analysis interface {
	CreateAnalysis(userId int, analysis models.AnalysisResponse) (string, error)
	GetAllAnalysis(id string) ([]models.Analysis, error)
	GetAnalysisChatHistory(analysisId string, userId int) ([]models.ChatMessage, error)
	AddChatMessage(analysisId string, userId int, sender, message string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     NewAuthRepository(db),
		User:     NewUserRepository(db),
		Analysis: NewAnalysisRepository(db),
	}
}
