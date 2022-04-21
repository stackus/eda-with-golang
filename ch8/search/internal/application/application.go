package application

import (
	"context"

	"eda-in-golang/ch8/search/internal/models"
)

type (
	SearchOrders struct{}

	GetOrder struct {
		OrderID string
	}

	Application interface {
		SearchOrders(ctx context.Context, search SearchOrders) ([]*models.Order, error)
		GetOrder(ctx context.Context, get GetOrder) (*models.Order, error)
	}

	app struct {
		orders OrderRepository
	}
)

var _ Application = (*app)(nil)

func New(orders OrderRepository) *app {
	return &app{
		orders: orders,
	}
}

func (a app) SearchOrders(ctx context.Context, search SearchOrders) ([]*models.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (a app) GetOrder(ctx context.Context, get GetOrder) (*models.Order, error) {
	// TODO implement me
	panic("implement me")
}
