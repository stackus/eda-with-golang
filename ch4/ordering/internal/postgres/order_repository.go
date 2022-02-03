package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type OrderRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(tableName string, db *sql.DB) OrderRepository {
	return OrderRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r OrderRepository) Find(ctx context.Context, orderID string) (*domain.Order, error) {
	const query = "SELECT items, card_token, sms_number, status FROM %s WHERE id = $1 LIMIT 1"

	order := &domain.Order{
		ID: orderID,
	}

	var items []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(&items, &order.CardToken, &order.SmsNumber, &status)
	if err != nil {
		return nil, errors.Wrap(err, "scanning order")
	}

	order.Status, err = r.statusToDomain(status)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(items, &order.Items)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling items")
	}

	return order, nil
}

func (r OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	const query = "INSERT INTO %s (id, items, card_token, sms_number, status) VALUES ($1, $2, $3, $4, $5)"

	items, err := json.Marshal(order.Items)
	if err != nil {
		return errors.Wrap(err, "marshalling items")
	}

	_, err = r.db.ExecContext(ctx, r.table(query), order.ID, items, order.CardToken, order.SmsNumber, order.Status.String())
	if err != nil {
		return errors.Wrap(err, "inserting order")
	}

	return nil
}

func (r OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	const query = "UPDATE %s SET items = $1, card_token = $2, sms_number = $3, status = $4 WHERE id = $5"

	items, err := json.Marshal(order.Items)
	if err != nil {
		return errors.Wrap(err, "marshalling items")
	}

	_, err = r.db.ExecContext(ctx, r.table(query), items, order.CardToken, order.SmsNumber, order.Status.String(), order.ID)
	if err != nil {
		return errors.Wrap(err, "updating order")
	}

	return nil
}

func (r OrderRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

func (r OrderRepository) statusToDomain(status string) (domain.OrderStatus, error) {
	switch status {
	case domain.OrderPending.String():
		return domain.OrderPending, nil
	case domain.OrderCancelled.String():
		return domain.OrderCancelled, nil
	case domain.OrderCompleted.String():
		return domain.OrderCompleted, nil
	default:
		return domain.OrderUnknown, fmt.Errorf("unknown order status: %s", status)
	}
}
