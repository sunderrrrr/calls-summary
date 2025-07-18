package repository

import (
	"api/pkg/utils/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Hostname string
	Port     string
	Username string
	Password string
	Dbname   string
	SSLMode  string
}

const (
	userDb    = "users"
	recordsDb = "calls"
)

func NewDB(config DB) (*sqlx.DB, error) {
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
