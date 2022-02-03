package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type CheckoutBasket struct {
	ID        string
	CardToken string
	SmsNumber string
}

type CheckoutBasketHandler struct {
	basketRepo domain.BasketRepository
	orderRepo  domain.OrderRepository
}

func NewCheckoutBasketHandler(basketRepo domain.BasketRepository, orderRepo domain.OrderRepository) CheckoutBasketHandler {
	return CheckoutBasketHandler{
		basketRepo: basketRepo,
		orderRepo:  orderRepo,
	}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.basketRepo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.Checkout(cmd.CardToken, cmd.SmsNumber)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	// submit the basket to the order module
	_, err = h.orderRepo.Save(ctx, basket)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	return errors.Wrap(h.basketRepo.Update(ctx, basket), "basket checkout")
}
