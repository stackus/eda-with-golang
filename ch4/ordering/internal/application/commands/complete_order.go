package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CompleteOrder struct {
	ID string
}

type CompleteOrderHandler struct {
	orders domain.OrderRepository
}

func NewCompleteOrderHandler(orders domain.OrderRepository) CompleteOrderHandler {
	return CompleteOrderHandler{
		orders: orders,
	}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = order.Complete()
	if err != nil {
		return nil
	}

	return h.orders.Update(ctx, order)
}
