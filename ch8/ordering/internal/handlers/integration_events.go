package handlers

import (
	"eda-in-golang/internal/ddd"
	"eda-in-golang/ordering/internal/domain"
)

func RegisterIntegrationEventHandlers[T ddd.AggregateEvent](eventHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.OrderCreatedEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
		domain.OrderCompletedEvent,
	)
}
