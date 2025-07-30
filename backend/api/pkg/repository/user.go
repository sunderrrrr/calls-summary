package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB // Подключение к базе данных
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	// Создает новый экземпляр репозитория пользователей
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) ResetPassword(email, newPassword string) error {
	// Обновляет пароль пользователя в базе данных
	// Возвращает ошибку, если операция не удалась
	query := fmt.Sprintf(`UPDATE %s SET pass_hash = $1 WHERE email = $2`, userDb)
	_, err := r.db.Exec(query, newPassword, email)
	if err != nil {
		return err
	}
	return nil
}
