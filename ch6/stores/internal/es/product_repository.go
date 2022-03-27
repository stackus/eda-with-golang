package es

import (
	"context"

	"eda-in-golang/ch6/internal/es"
	"eda-in-golang/ch6/internal/registry"
	"eda-in-golang/ch6/stores/internal/domain"
)

type ProductRepository struct {
	aggregates es.AggregateRepository
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(registry registry.Registry, store es.AggregateStore) ProductRepository {
	return ProductRepository{
		aggregates: es.NewAggregateRepository(registry, store),
	}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	agg, err := r.aggregates.Load(ctx, productID, domain.ProductAggregate)
	if err != nil {
		return nil, err
	}

	return agg.(*domain.Product), nil
}

func (r ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	return r.aggregates.Save(ctx, product)
}
