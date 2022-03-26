package application

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/ordering/internal/domain"
)

type NotificationHandlers struct {
	notifications domain.NotificationRepository
}

var _ ddd.EventHandler = (*NotificationHandlers)(nil)

func NewNotificationHandlers(notifications domain.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notifications: notifications,
	}
}

func (h NotificationHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case domain.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}
	return nil
}

func (h NotificationHandlers) onOrderCreated(ctx context.Context, event ddd.Event) error {
	orderCreated := event.Payload().(*domain.OrderCreated)
	return h.notifications.NotifyOrderCreated(ctx, orderCreated.Order.ID(), orderCreated.Order.CustomerID)
}

func (h NotificationHandlers) onOrderReadied(ctx context.Context, event ddd.Event) error {
	orderReadied := event.Payload().(*domain.OrderReadied)
	return h.notifications.NotifyOrderReady(ctx, orderReadied.Order.ID(), orderReadied.Order.CustomerID)
}

func (h NotificationHandlers) onOrderCanceled(ctx context.Context, event ddd.Event) error {
	orderCanceled := event.Payload().(*domain.OrderCanceled)
	return h.notifications.NotifyOrderCanceled(ctx, orderCanceled.Order.ID(), orderCanceled.Order.CustomerID)
}
