package domain

import (
	"fmt"
)

var (
	ErrOrderHasNoItems        = fmt.Errorf("the order has no items")
	ErrCardTokenCannotBeBlank = fmt.Errorf("the card token cannot be blank")
	ErrSmsNumberCannotBeBlank = fmt.Errorf("the SMS number cannot be blank")
)

type Order struct {
	ID        string
	Items     []*Item
	CardToken string
	SmsNumber string
	Status    OrderStatus
}

func CreateOrder(id string, items []*Item, cardToken, smsNumber string) (*Order, error) {
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
