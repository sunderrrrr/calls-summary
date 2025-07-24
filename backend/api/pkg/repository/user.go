package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) ResetPassword(email, newPassword string) error {
	query := fmt.Sprintf(`UPDATE %s SET pass_hash = $1 WHERE email = $2`, userDb)
	_, err := r.db.Exec(query, newPassword, email)
	if err != nil {
		return err
	}
	return nil
}
