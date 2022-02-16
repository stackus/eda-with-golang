package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CancelOrder struct {
	ID domain.OrderID
}

type CancelOrderHandler struct {
	orders        domain.OrderRepository
	invoices      domain.InvoiceRepository
	shoppingLists domain.ShoppingListRepository
}

func NewCancelOrderHandler(orders domain.OrderRepository, invoices domain.InvoiceRepository, shoppingLists domain.ShoppingListRepository) CancelOrderHandler {
	return CancelOrderHandler{
		orders:        orders,
		invoices:      invoices,
		shoppingLists: shoppingLists,
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

	err = h.invoices.Delete(ctx, order.InvoiceID)
	if err != nil {
		return err
	}

	err = h.shoppingLists.Delete(ctx, order.ID)
	if err != nil {
		return err
	}

	return h.orders.Update(ctx, order)
}
