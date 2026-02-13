package service

import (
	"time"

	subsapp "github.com/VSBrilyakov/subsApp"
	"github.com/VSBrilyakov/subsApp/internal/repository"
)

type SubscriptionService struct {
	repo repository.SubscribeActions
}

func NewSubscriptionService(repo *repository.Repository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(sub subsapp.Subscription) (int, error) {
	return s.repo.CreateSubscription(sub)
}

func (s *SubscriptionService) GetSubscription(subId int) (*subsapp.Subscription, error) {
	return s.repo.GetSubscription(subId)
}

func (s *SubscriptionService) UpdateSubscription(subId int, input subsapp.UpdSubscription) error {
	return s.repo.UpdateSubscription(subId, input)
}

func (s *SubscriptionService) DeleteSubscription(subId int) error {
	return s.repo.DeleteSubscription(subId)
}

func (s *SubscriptionService) GetAllSubscriptions() (*[]subsapp.Subscription, error) {
	return s.repo.GetAllSubscriptions()
}

func (s *SubscriptionService) GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error) {
	return s.repo.GetSubsSumByUserID(userId, serviceName, dateFrom, dateTo)
}
