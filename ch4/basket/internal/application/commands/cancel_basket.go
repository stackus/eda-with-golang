package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
)

type CancelBasket struct {
	ID string
}

type CancelBasketHandler struct {
	repo ports.BasketRepository
}

func NewCancelBasketHandler(repo ports.BasketRepository) CancelBasketHandler {
	return CancelBasketHandler{repo: repo}
}

func (h CancelBasketHandler) CancelBasket(ctx context.Context, cmd CancelBasket) error {
	basket, err := h.repo.FindBasket(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.Cancel()
	if err != nil {
		return err
	}

	return h.repo.UpdateBasket(ctx, basket)
}
