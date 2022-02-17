package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CancelOrder struct {
	ID domain.OrderID
}

type CancelOrderHandler struct {
	orders   domain.OrderRepository
	invoices domain.InvoiceRepository
	shopping domain.ShoppingRepository
}

func NewCancelOrderHandler(orders domain.OrderRepository, shopping domain.ShoppingRepository) CancelOrderHandler {
	return CancelOrderHandler{
		orders:   orders,
		shopping: shopping,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = order.Cancel()
	if err != nil {
		return err
	}

	err = h.shopping.Cancel(ctx, order.ShoppingID)
	if err != nil {
		return err
	}

	return h.orders.Update(ctx, order)
}
