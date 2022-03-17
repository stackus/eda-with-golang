package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/baskets/internal/domain"
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
	case domain.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	}
	return nil
}

func (h OrderHandlers) onBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	checkedOut := event.Payload().(*domain.BasketCheckedOut)
	_, err := h.orders.Save(ctx, checkedOut.Basket)
	return err
}
