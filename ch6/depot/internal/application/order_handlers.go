package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type OrderHandlers struct {
	orders domain.OrderRepository
}

var _ ddd.EventHandler = (*OrderHandlers)(nil)

func NewOrderHandlers(orders domain.OrderRepository) OrderHandlers {
	return OrderHandlers{
		orders: orders,
	}
}

func (h OrderHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}
	return nil
}

func (h OrderHandlers) onShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	completed := event.Payload().(*domain.ShoppingListCompleted)
	return h.orders.Ready(ctx, completed.ShoppingList.OrderID)
}
