package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) signUp(email, password string) (int, error) {
	// mock
	return 1, nil
}

func (r *AuthRepository) signIn(email, password string) (models.User, error) {
	return models.User{}, nil
}
