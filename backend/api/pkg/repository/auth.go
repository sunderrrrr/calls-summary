package repository

import (
	"api/models"
	"fmt"
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

func (r *AuthRepository) SignUp(email, name, passwordHash string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, pass_hash) VALUES ($1, $2, $3) RETURNING id", userDb)
	if err := r.db.QueryRow(query, email, name, passwordHash).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, name, email FROM %s WHERE email = $1 AND pass_hash = $2", userDb)
	if err := r.db.Get(&user, query, email, password); err != nil {
		return models.User{}, err
	}
	return user, nil
}
