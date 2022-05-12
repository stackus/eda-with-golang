package application

import (
	"context"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	// case domain.StoreCreatedEvent:
	// 	return h.onStoreCreated(ctx, event)
	// case domain.StoreParticipationEnabledEvent:
	// 	return h.onStoreParticipationEnabled(ctx, event)
	// case domain.StoreParticipationDisabledEvent:
	// 	return h.onStoreParticipationDisabled(ctx, event)
	// case domain.StoreRebrandedEvent:
	// 	return h.onStoreRebranded(ctx, event)
	}
	return nil
}

// func (h IntegrationEventHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
// 	payload := event.Payload().(*domain.StoreRebranded)
// 	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
// 		ddd.NewEvent(storespb.StoreRebrandedEvent, &storespb.StoreRebranded{
// 			Id:   event.AggregateID(),
// 			Name: payload.Name,
// 		}),
// 	)
// }
