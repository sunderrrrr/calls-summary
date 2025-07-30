package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB // Подключение к базе данных
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	// Создает новый экземпляр репозитория авторизации
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) SignUp(email, name, passwordHash string) (int, error) {
	// Регистрирует нового пользователя в базе данных
	// Возвращает ID пользователя или ошибку
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, pass_hash) VALUES ($1, $2, $3) RETURNING id", userDb)
	if err := r.db.QueryRow(query, email, name, passwordHash).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(email, password string) (models.User, error) {
	// Получает пользователя из базы данных по email и паролю
	// Возвращает объект пользователя или ошибку
	var user models.User
	query := fmt.Sprintf("SELECT id, name, email FROM %s WHERE email = $1 AND pass_hash = $2", userDb)
	if err := r.db.Get(&user, query, email, password); err != nil {
		return models.User{}, err
	}
	return user, nil
}
