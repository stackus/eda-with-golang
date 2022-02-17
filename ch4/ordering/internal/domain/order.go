package domain

import (
	"fmt"
)

var (
	ErrOrderHasNoItems         = fmt.Errorf("the order has no items")
	ErrOrderCannotBeCancelled  = fmt.Errorf("the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = fmt.Errorf("the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = fmt.Errorf("the payment id cannot be blank")
)

type OrderID string

type Order struct {
	ID         OrderID
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []*Item
	Status     OrderStatus
}

func (i OrderID) String() string {
	return string(i)
}

func ToOrderID(id string) OrderID {
	return OrderID(id)
}

func CreateOrder(id OrderID, customerID, paymentID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := &Order{
		ID:         id,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
		Status:     OrderPending,
	}

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderCancelled

	return nil
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price
	}

	return total
}
