package application

import (
	"context"

	"eda-in-golang/ch7/baskets/internal/domain"
	"eda-in-golang/ch7/internal/ddd"
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
	_, err := h.orders.Save(ctx, checkedOut.PaymentID, checkedOut.CustomerID, checkedOut.Items)
	return err
}
