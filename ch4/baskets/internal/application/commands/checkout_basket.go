package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type CheckoutBasket struct {
	ID        domain.BasketID
	CardToken string
	SmsNumber string
}

type CheckoutBasketHandler struct {
	baskets domain.BasketRepository
	orders  domain.OrderRepository
}

func NewCheckoutBasketHandler(baskets domain.BasketRepository, orders domain.OrderRepository) CheckoutBasketHandler {
	return CheckoutBasketHandler{
		baskets: baskets,
		orders:  orders,
	}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.baskets.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.Checkout(cmd.CardToken, cmd.SmsNumber)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	// submit the basket to the order module
	_, err = h.orders.Save(ctx, basket)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	return errors.Wrap(h.baskets.Update(ctx, basket), "basket checkout")
}
