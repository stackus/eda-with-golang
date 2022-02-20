package handlers

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/ddd"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type NotificationHandlers struct {
	notifications domain.NotificationRepository
}

func RegisterNotificationHandlers(domainSubscriber ddd.EventSubscriber, notifications domain.NotificationRepository) {
	handlers := NotificationHandlers{
		notifications: notifications,
	}

	domainSubscriber.Subscribe(domain.OrderCreated{}, handlers.OnOrderCreated)
	domainSubscriber.Subscribe(domain.OrderReadied{}, handlers.OnOrderReadied)
	domainSubscriber.Subscribe(domain.OrderCanceled{}, handlers.OnOrderCanceled)
}

func (h NotificationHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	orderCreated := event.(*domain.OrderCreated)
	return h.notifications.NotifyOrderCreated(ctx, orderCreated.Order.ID, orderCreated.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	orderReadied := event.(*domain.OrderReadied)
	return h.notifications.NotifyOrderReady(ctx, orderReadied.Order.ID, orderReadied.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderCanceled(ctx context.Context, event ddd.Event) error {
	orderCanceled := event.(*domain.OrderCanceled)
	return h.notifications.NotifyOrderCanceled(ctx, orderCanceled.Order.ID, orderCanceled.Order.CustomerID)
}
