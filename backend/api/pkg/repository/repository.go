package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
	User
}
type Auth interface {
	SignUp(email, name, password string) (int, error)
	GetUser(email, password string) (models.User, error)
}
type User interface {
	ResetPassword(email, newPassword string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
		User: NewUserRepository(db),
	}
}
