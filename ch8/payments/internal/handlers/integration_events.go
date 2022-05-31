package handlers

import (
	"eda-in-golang/internal/ddd"
	"eda-in-golang/payments/internal/models"
)

func RegisterIntegrationEventHandlers[T ddd.Event](eventHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(eventHandlers,
		models.InvoicePaidEvent,
	)
}
