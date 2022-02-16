package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type CancelBasket struct {
	ID domain.BasketID
}

type CancelBasketHandler struct {
	repo domain.BasketRepository
}

func NewCancelBasketHandler(repo domain.BasketRepository) CancelBasketHandler {
	return CancelBasketHandler{repo: repo}
}

func (h CancelBasketHandler) CancelBasket(ctx context.Context, cmd CancelBasket) error {
	basket, err := h.repo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.Cancel()
	if err != nil {
		return err
	}

	return h.repo.Update(ctx, basket)
}
