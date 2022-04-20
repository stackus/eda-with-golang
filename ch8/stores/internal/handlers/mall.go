package handlers

import (
	"eda-in-golang/ch8/internal/ddd"
	"eda-in-golang/ch8/stores/internal/domain"
)

func RegisterMallHandlers[T ddd.AggregateEvent](mallHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(mallHandlers,
		domain.StoreCreatedEvent,
		domain.StoreParticipationEnabledEvent,
		domain.StoreParticipationDisabledEvent,
		domain.StoreRebrandedEvent,
	)
}
