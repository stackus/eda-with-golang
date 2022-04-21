package application

import (
	"context"

	"eda-in-golang/ch8/search/internal/models"
)

type OrderRepository interface {
	Add(ctx context.Context) error
	Search(ctx context.Context) ([]*models.Order, error)
	Get(ctx context.Context, orderID string) (*models.Order, error)
}
