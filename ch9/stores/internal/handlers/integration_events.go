package handlers

import (
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/stores/internal/domain"
)

func RegisterIntegrationEventHandlers[T ddd.AggregateEvent](eventHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.StoreCreatedEvent,
		domain.StoreParticipationEnabledEvent,
		domain.StoreParticipationDisabledEvent,
		domain.StoreRebrandedEvent,
	)
}
