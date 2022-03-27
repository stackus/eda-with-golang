package es

import (
	"context"

	"eda-in-golang/ch7/internal/es"
	"eda-in-golang/ch7/internal/registry"
	"eda-in-golang/ch7/ordering/internal/domain"
)

type OrderRepository struct {
	aggregates es.AggregateRepository
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(registry registry.Registry, order es.AggregateStore) OrderRepository {
	return OrderRepository{
		aggregates: es.NewAggregateRepository(registry, order),
	}
}

func (r OrderRepository) Find(ctx context.Context, orderID string) (*domain.Order, error) {
	agg, err := r.aggregates.Load(ctx, orderID, domain.OrderAggregate)
	if err != nil {
		return nil, err
	}

	return agg.(*domain.Order), nil
}

func (r OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	return r.aggregates.Save(ctx, order)
}
