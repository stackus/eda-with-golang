package domain

import (
	"fmt"
)

var (
	ErrBasketCannotBeCancelled  = fmt.Errorf("the basket cannot be cancelled")
	ErrBasketCannotBeModified   = fmt.Errorf("the basket cannot be modified")
	ErrQuantityCannotBeNegative = fmt.Errorf("the item quantity cannot be negative")
	ErrCardTokenCannotBeBlank   = fmt.Errorf("the card token cannot be blank")
	ErrSmsNumberCannotBeBlank   = fmt.Errorf("the sms number cannot be blank")
)

type BasketStatus string

const (
	BasketUnknown    BasketStatus = ""
	BasketOpen       BasketStatus = "open"
	BasketCancelled  BasketStatus = "cancelled"
	BasketCheckedOut BasketStatus = "checked_out"
)

func (s BasketStatus) String() string {
	switch s {
	case BasketOpen, BasketCancelled, BasketCheckedOut:
		return string(s)
	default:
		return ""
	}
}

type Basket struct {
	ID        string
	Items     []Item
	CardToken string
	SmsNumber string
	Status    BasketStatus
}

func StartBasket(id string) (basket *Basket) {
	basket = &Basket{
		ID:     id,
		Status: BasketOpen,
		Items:  []Item{},
	}

	return
}

func (b Basket) IsCancellable() bool {
	return b.Status == BasketOpen
}

func (b Basket) IsOpen() bool {
	return b.Status == BasketOpen
}

func (b *Basket) Cancel() error {
	if !b.IsCancellable() {
		return ErrBasketCannotBeCancelled
	}

	b.Status = BasketCancelled
	b.Items = []Item{}

	return nil
}

func (b *Basket) Checkout(cardToken, smsNumber string) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if cardToken == "" {
		return ErrCardTokenCannotBeBlank
	}

	if smsNumber == "" {
		return ErrSmsNumberCannotBeBlank
	}

	b.CardToken = cardToken
	b.SmsNumber = smsNumber
	b.Status = BasketCheckedOut

	return nil
}

func (b *Basket) AddItem(product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity += quantity
			return nil
		}
	}

	b.Items = append(b.Items, Item{
		StoreID:      product.StoreID,
		ProductID:    product.ID,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     quantity,
	})

	return nil
}

func (b *Basket) RemoveItem(product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity -= quantity

			if b.Items[i].Quantity < 1 {
				b.Items = append(b.Items[:i], b.Items[i+1:]...)
			}
			return nil
		}
	}

	return nil
}
