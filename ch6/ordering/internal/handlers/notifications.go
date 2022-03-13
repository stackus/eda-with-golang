package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/ordering/internal/application"
	"github.com/stackus/eda-with-golang/ch6/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers application.DomainEventHandlers,
	domainSubscriber ddd.EventSubscriber,
) {
	domainSubscriber.Subscribe(domain.OrderCreated{}, notificationHandlers.OnOrderCreated)
	domainSubscriber.Subscribe(domain.OrderReadied{}, notificationHandlers.OnOrderReadied)
	domainSubscriber.Subscribe(domain.OrderCanceled{}, notificationHandlers.OnOrderCanceled)
}
