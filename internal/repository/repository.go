package repository

import "github.com/jmoiron/sqlx"

type SubscribeActions interface {
}
type Repository struct {
	SubscribeActions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
