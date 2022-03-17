package es

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type AggregateRepository struct {
	registry registry.Registry
	store    AggregateStore
}

func NewAggregateRepository(registry registry.Registry, store AggregateStore) AggregateRepository {
	return AggregateRepository{
		registry: registry,
		store:    store,
	}
}

func (r AggregateRepository) Load(ctx context.Context, aggregateID, aggregateName string) (interface{}, error) {
	agg, err := r.registry.Build(
		aggregateName,
		ddd.SetID(aggregateID),
		ddd.SetName(aggregateName),
	)
	if err != nil {
		return nil, err
	}

	if err = r.store.Load(ctx, agg.(EventSourcedAggregate)); err != nil {
		return nil, err
	}

	return agg, nil
}

func (r AggregateRepository) Save(ctx context.Context, aggregate EventSourcedAggregate) error {
	if aggregate.Version() == aggregate.PendingVersion() {
		return nil
	}

	for _, event := range aggregate.Events() {
		if err := aggregate.ApplyEvent(event); err != nil {
			return err
		}
	}

	err := r.store.Save(ctx, aggregate)
	if err != nil {
		return err
	}

	aggregate.CommitEvents()

	return nil
}
