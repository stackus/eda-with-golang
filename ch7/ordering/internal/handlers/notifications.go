package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/ordering/internal/domain"
)

func RegisterNotificationHandlers[T ddd.AggregateEvent](notificationHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(domain.OrderCreatedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderCanceledEvent, notificationHandlers)
}
