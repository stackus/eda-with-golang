package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/ordering/internal/domain"
)

type NotificationHandlers struct {
	notifications domain.NotificationRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*NotificationHandlers)(nil)

func NewNotificationHandlers(notifications domain.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notifications: notifications,
	}
}

func (h NotificationHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	orderCreated := event.Payload().(*domain.OrderCreated)
	return h.notifications.NotifyOrderCreated(ctx, orderCreated.Order.ID(), orderCreated.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	orderReadied := event.Payload().(*domain.OrderReadied)
	return h.notifications.NotifyOrderReady(ctx, orderReadied.Order.ID(), orderReadied.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderCanceled(ctx context.Context, event ddd.Event) error {
	orderCanceled := event.Payload().(*domain.OrderCanceled)
	return h.notifications.NotifyOrderCanceled(ctx, orderCanceled.Order.ID(), orderCanceled.Order.CustomerID)
}