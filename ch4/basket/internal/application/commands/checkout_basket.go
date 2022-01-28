package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
)

type CheckoutBasket struct {
	ID        string
	CardToken string
	SmsNumber string
}

type CheckoutBasketHandler struct {
	repo ports.BasketRepository
}

func NewCheckoutBasketHandler(repo ports.BasketRepository) CheckoutBasketHandler {
	return CheckoutBasketHandler{repo: repo}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.repo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.Checkout(cmd.CardToken, cmd.SmsNumber)
	if err != nil {
		return err
	}

	return h.repo.Update(ctx, basket)
}
