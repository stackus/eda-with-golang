package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"eda-in-golang/ch8/search/internal/application"
	"eda-in-golang/ch8/search/internal/models"
)

type OrderRepository struct {
	tableName string
	db        *sql.DB
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(tableName string, db *sql.DB) OrderRepository {
	return OrderRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r OrderRepository) Add(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (r OrderRepository) Search(ctx context.Context) ([]*models.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r OrderRepository) Get(ctx context.Context, orderID string) (*models.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r OrderRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
