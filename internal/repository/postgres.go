package repository

import (
	"fmt"

	"githhub.com/VSBrilyakov/test-app/configs"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *configs.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
