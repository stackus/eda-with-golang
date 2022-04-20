package handlers

import (
	"eda-in-golang/ch8/internal/ddd"
	"eda-in-golang/ch8/payments/internal/models"
)

func RegisterIntegrationEventHandlers[T ddd.Event](eventHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(eventHandlers,
		models.InvoicePaidEvent,
	)
}
