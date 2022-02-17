package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CreateOrder struct {
	ID         domain.OrderID
	CustomerID string
	PaymentID  string
	Items      []*domain.Item
}

type CreateOrderHandler struct {
	orders    domain.OrderRepository
	customers domain.CustomerRepository
	payments  domain.PaymentRepository
	shopping  domain.ShoppingRepository
}

func NewCreateOrderHandler(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository, shopping domain.ShoppingRepository) CreateOrderHandler {
	return CreateOrderHandler{
		orders:    orders,
		customers: customers,
		payments:  payments,
		shopping:  shopping,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if err = h.customers.Authorize(ctx, order.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	// validatePayment
	if err = h.payments.Confirm(ctx, order.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	// scheduleShopping
	if shoppingID, err := h.shopping.Create(ctx, order); err != nil {
		return errors.Wrap(err, "order shopping scheduling")
	} else {
		order.ShoppingID = shoppingID
	}

	// inProgressOrder

	return errors.Wrap(h.orders.Save(ctx, order), "create order command")
}
