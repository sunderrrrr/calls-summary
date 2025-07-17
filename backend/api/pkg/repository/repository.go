package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
}
type Auth interface {
	signUp(email, password string) (int, error)
	signIn(email, password string) (models.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
