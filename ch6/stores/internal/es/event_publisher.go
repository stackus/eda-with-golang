package es

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/es"
)

type EventPublisher struct {
	es.AggregateStore
	publisher ddd.EventPublisher
}

var _ es.AggregateStore = (*EventPublisher)(nil)

func NewEventPublisher(store es.AggregateStore, publisher ddd.EventPublisher) EventPublisher {
	return EventPublisher{
		AggregateStore: store,
		publisher:      publisher,
	}
}

func (p EventPublisher) Save(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	if err := p.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}
	return p.publisher.Publish(ctx, aggregate.GetEvents()...)
}
