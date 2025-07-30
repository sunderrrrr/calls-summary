package repository

import (
	"api/pkg/utils/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Hostname string // Хост базы данных
	Port     string // Порт базы данных
	Username string // Имя пользователя базы данных
	Password string // Пароль пользователя базы данных
	Dbname   string // Имя базы данных
	SSLMode  string // Режим SSL подключения
}

const (
	userDb    = "users"
	recordsDb = "calls"
)

func NewDB(config DB) (*sqlx.DB, error) {
	// Создает новое подключение к базе данных с использованием переданной конфигурации
	// Проверяет соединение с базой данных
	// Возвращает объект подключения или ошибку
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Hostname, config.Port, config.Username, config.Password, config.Dbname, config.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}
	return db, nil
}
