package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CreateOrder struct {
	ID        domain.OrderID
	Items     []*domain.Item
	CardToken string
	SmsNumber string
}

type CreateOrderHandler struct {
	orders        domain.OrderRepository
	invoices      domain.InvoiceRepository
	shoppingLists domain.ShoppingListRepository
}

func NewCreateOrderHandler(orders domain.OrderRepository, invoices domain.InvoiceRepository, shoppingLists domain.ShoppingListRepository) CreateOrderHandler {
	return CreateOrderHandler{
		orders:        orders,
		invoices:      invoices,
		shoppingLists: shoppingLists,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.Items, cmd.CardToken, cmd.SmsNumber)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	order.InvoiceID, err = h.invoices.Save(ctx, order.ID, order.GetTotal())
	if err != nil {
		return err
	}

	err = h.shoppingLists.Save(ctx, order)
	if err != nil {
		errors.Wrap(err, "create order command")
	}

	return errors.Wrap(h.orders.Save(ctx, order), "create order command")
}
