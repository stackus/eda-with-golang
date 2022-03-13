package es

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/es"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type StoreRepository struct {
	aggregates es.AggregateRepository
}

var _ domain.StoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(registry registry.Registry, store es.AggregateStore) StoreRepository {
	return StoreRepository{
		aggregates: es.NewAggregateRepository(registry, store),
	}
}

func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	agg, err := r.aggregates.Load(ctx, storeID, domain.StoreAggregate)
	if err != nil {
		return nil, err
	}

	return agg.(*domain.Store), nil
}

func (r StoreRepository) Save(ctx context.Context, store *domain.Store) error {
	return r.aggregates.Save(ctx, store)
}
