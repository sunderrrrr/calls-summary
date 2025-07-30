package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth     // Интерфейс для работы с авторизацией
	User     // Интерфейс для работы с пользователями
	Analysis // Интерфейс для работы с анализами и чатом
}

type Auth interface {
	SignUp(email, name, password string) (int, error)    // Регистрация нового пользователя
	GetUser(email, password string) (models.User, error) // Получение пользователя по email и паролю
}

type User interface {
	ResetPassword(email, newPassword string) error // Сброс пароля пользователя
}

type Analysis interface {
	CreateAnalysis(userId int, analysis models.AnalysisResponse) (string, error)        // Создание нового анализа
	GetAllAnalysis(id string) ([]models.Analysis, error)                                // Получение всех анализов пользователя
	GetAnalysisChatHistory(analysisId string, userId int) ([]models.ChatMessage, error) // Получение истории чата анализа
	AddChatMessage(analysisId string, userId int, sender, message string) error         // Добавление сообщения в чат
}

func NewRepository(db *sqlx.DB) *Repository {
	// Создает новый экземпляр репозитория с подключением к базе данных
	return &Repository{
		Auth:     NewAuthRepository(db),
		User:     NewUserRepository(db),
		Analysis: NewAnalysisRepository(db),
	}
}
