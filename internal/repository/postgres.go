package repository

import (
	"fmt"

	"github.com/VSBrilyakov/subsApp/configs"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

const (
	subscriptionTable = "subscription"
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

func DoMigrates(db *sqlx.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db.DB, "./migrations"); err != nil {
		return err
	}

	return nil
}
