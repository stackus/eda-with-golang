package handlers

import (
	"eda-in-golang/ch8/customers/internal/domain"
	"eda-in-golang/ch8/internal/ddd"
)

func RegisterIntegrationEventHandlers[T ddd.AggregateEvent](eventHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.CustomerRegisteredEvent,
		domain.CustomerSmsChangedEvent,
		domain.CustomerEnabledEvent,
		domain.CustomerDisabledEvent,
	)
}
