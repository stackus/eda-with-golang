package es

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type EventPublisher struct {
	AggregateStore
	publisher ddd.EventPublisher
}

var _ AggregateStore = (*EventPublisher)(nil)

func NewEventPublisher(store AggregateStore, publisher ddd.EventPublisher) EventPublisher {
	return EventPublisher{
		AggregateStore: store,
		publisher:      publisher,
	}
}

func (p EventPublisher) Save(ctx context.Context, aggregate EventSourcedAggregate) error {
	if err := p.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}
	return p.publisher.Publish(ctx, aggregate.GetEvents()...)
}
