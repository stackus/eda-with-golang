package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCreatedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderCanceledEvent, notificationHandlers)
}
