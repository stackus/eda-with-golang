package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch5/internal/ddd"
)

type OrderHandlers struct {
	orders domain.OrderRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*OrderHandlers)(nil)

func NewOrderHandlers(orders domain.OrderRepository) OrderHandlers {
	return OrderHandlers{
		orders: orders,
	}
}

func (h OrderHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	completed := event.(*domain.ShoppingListCompleted)
	return h.orders.Ready(ctx, completed.ShoppingList.OrderID)
}
