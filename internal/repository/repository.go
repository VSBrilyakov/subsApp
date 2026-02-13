package repository

import (
	"time"

	subsapp "github.com/VSBrilyakov/subsApp"
	"github.com/jmoiron/sqlx"
)

type SubscribeActions interface {
	CreateSubscription(sub subsapp.Subscription) (int, error)
	GetSubscription(subId int) (*subsapp.Subscription, error)
	UpdateSubscription(subId int, input subsapp.UpdSubscription) error
	DeleteSubscription(subId int) error
	GetAllSubscriptions() (*[]subsapp.Subscription, error)
	GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error)
}
type Repository struct {
	SubscribeActions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SubscribeActions: NewSubPostgres(db),
	}
}
