package es

import (
	"context"

	"eda-in-golang/ch6/baskets/internal/domain"
	"eda-in-golang/ch6/internal/es"
	"eda-in-golang/ch6/internal/registry"
)

type BasketRepository struct {
	aggregates es.AggregateRepository
}

var _ domain.BasketRepository = (*BasketRepository)(nil)

func NewBasketRepository(registry registry.Registry, basket es.AggregateStore) BasketRepository {
	return BasketRepository{
		aggregates: es.NewAggregateRepository(registry, basket),
	}
}

func (r BasketRepository) Find(ctx context.Context, basketID string) (*domain.Basket, error) {
	agg, err := r.aggregates.Load(ctx, basketID, domain.BasketAggregate)
	if err != nil {
		return nil, err
	}

	return agg.(*domain.Basket), nil
}

func (r BasketRepository) Save(ctx context.Context, basket *domain.Basket) error {
	return r.aggregates.Save(ctx, basket)
}
