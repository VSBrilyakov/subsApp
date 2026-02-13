package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	subsapp "github.com/VSBrilyakov/subsApp"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SubPostgres struct {
	db *sqlx.DB
}

func NewSubPostgres(db *sqlx.DB) *SubPostgres {
	return &SubPostgres{db: db}
}

func (s *SubPostgres) CreateSubscription(sub subsapp.Subscription) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id", subscriptionTable)
	logrus.Debug(fmt.Sprintf("CreateSubscription query: %s", query))

	row := s.db.QueryRow(query, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SubPostgres) GetSubscription(subId int) (*subsapp.Subscription, error) {
	var sub subsapp.Subscription
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", subscriptionTable)
	logrus.Debug(fmt.Sprintf("GetSubscription query: %s", query))

	err := s.db.Get(&sub, query, subId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("subscription not found")
	}

	return &sub, err
}

func (s *SubPostgres) UpdateSubscription(subId int, input subsapp.UpdSubscription) error {
	if _, err := s.GetSubscription(subId); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ServiceName != nil {
		setValues = append(setValues, fmt.Sprintf("service_name=$%d", argId))
		args = append(args, *input.ServiceName)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.UserID != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserID)
		argId++
	}

	if input.StartDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, *input.StartDate)
		argId++
	}

	if input.EndDate != nil {
		setValues = append(setValues, fmt.Sprintf("end_date=$%d", argId))
		args = append(args, *input.EndDate)
		argId++
	}

	if argId == 1 {
		return errors.New("empty fields")
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", subscriptionTable, setQuery, argId)
	args = append(args, subId)

	logrus.Debug(fmt.Sprintf("updateQuery: %s", query))
	logrus.Debug(fmt.Sprintf("args: %s", args))

	_, err := s.db.Exec(query, args...)
	return err
}

func (s *SubPostgres) DeleteSubscription(subId int) error {
	if _, err := s.GetSubscription(subId); err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", subscriptionTable)
	logrus.Debug(fmt.Sprintf("deleteQuery: %s", query))
	_, err := s.db.Exec(query, subId)

	return err
}

func (s *SubPostgres) GetAllSubscriptions() (*[]subsapp.Subscription, error) {
	var subs []subsapp.Subscription
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", subscriptionTable)
	logrus.Debug("GetAllSubscriptions query: %s", query)

	err := s.db.Select(&subs, query)
	if err != nil {
		return nil, err
	}

	return &subs, nil
}

func (s *SubPostgres) GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error) {
	var tp int
	query := fmt.Sprintf("SELECT COALESCE(SUM(price), 0)::INTEGER AS total_price FROM %s WHERE user_id = $1 AND service_name = $2 AND start_date BETWEEN $3 AND $4",
		subscriptionTable)
	logrus.Debug(fmt.Sprintf("GetSubsSumByUserID: %s", query))
	logrus.Debug(fmt.Sprintf("args: %s %s %s %s", userId, serviceName, dateFrom, dateTo))

	row := s.db.QueryRow(query, userId, serviceName, dateFrom, dateTo)
	if err := row.Scan(&tp); err != nil {
		return 0, err
	}

	return tp, nil
}
