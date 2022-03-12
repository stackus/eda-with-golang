package domain

import (
	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

const OrderAggregate = "ordering.Order"

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	ddd.Aggregate
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []*Item
	Status     OrderStatus
}

func NewOrder(id string) *Order {
	return &Order{
		Aggregate: ddd.NewAggregate(id, OrderAggregate),
	}
}

func CreateOrder(id, customerID, paymentID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := NewOrder(id)
	order.CustomerID = customerID
	order.PaymentID = paymentID
	order.Items = items
	order.Status = OrderIsPending

	order.AddEvent(OrderCreatedEvent, &OrderCreated{
		Order: order,
	})

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderIsPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderIsCancelled

	o.AddEvent(OrderCanceledEvent, &OrderCanceled{
		Order: o,
	})
	return nil
}

func (o *Order) Ready() error {
	// validate status

	o.Status = OrderIsReady

	o.AddEvent(OrderReadiedEvent, &OrderReadied{
		Order: o,
	})

	return nil
}

func (o *Order) Complete(invoiceID string) error {
	// validate invoice exists

	// validate status

	o.InvoiceID = invoiceID
	o.Status = OrderIsCompleted

	o.AddEvent(OrderCompletedEvent, &OrderCompleted{
		Order: o,
	})

	return nil
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
