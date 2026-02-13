package service

import (
	"time"

	subsapp "github.com/VSBrilyakov/subsApp"
	"github.com/VSBrilyakov/subsApp/internal/repository"
)

type SubscribeActions interface {
	CreateSubscription(sub subsapp.Subscription) (int, error)
	GetSubscription(subId int) (*subsapp.Subscription, error)
	UpdateSubscription(subId int, input subsapp.UpdSubscription) error
	DeleteSubscription(subId int) error
	GetAllSubscriptions() (*[]subsapp.Subscription, error)
	GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error)
}
type Service struct {
	SubscribeActions
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SubscribeActions: NewSubscriptionService(repo),
	}
}
