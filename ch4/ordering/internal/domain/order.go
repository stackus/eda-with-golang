package domain

import (
	"fmt"
)

var (
	ErrOrderHasNoItems        = fmt.Errorf("the order has no items")
	ErrOrderCannotBeCancelled = fmt.Errorf("the order cannot be cancelled")
	ErrCardTokenCannotBeBlank = fmt.Errorf("the card token cannot be blank")
	ErrSmsNumberCannotBeBlank = fmt.Errorf("the SMS number cannot be blank")
)

type OrderID string

type Order struct {
	ID        OrderID
	Items     []*Item
	CardToken string
	SmsNumber string
	InvoiceID InvoiceID
	Status    OrderStatus
}

func (i OrderID) String() string {
	return string(i)
}

func CreateOrder(id OrderID, items []*Item, cardToken, smsNumber string) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if cardToken == "" {
		return nil, ErrCardTokenCannotBeBlank
	}

	if smsNumber == "" {
		return nil, ErrSmsNumberCannotBeBlank
	}

	order := &Order{
		ID:        id,
		Items:     items,
		CardToken: cardToken,
		SmsNumber: smsNumber,
		Status:    OrderPending,
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

func (o *Order) GetInvoice() *Invoice {
	return &Invoice{
		ID:     o.InvoiceID,
		Amount: o.GetTotal(),
	}
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price
	}

	return total
}
