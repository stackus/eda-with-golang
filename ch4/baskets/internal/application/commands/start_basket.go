package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type StartBasket struct {
	ID string
}

type StartBasketHandler struct {
	repo domain.BasketRepository
}

func NewStartBasketHandler(repo domain.BasketRepository) StartBasketHandler {
	return StartBasketHandler{repo: repo}
}

func (h StartBasketHandler) StartBasket(ctx context.Context, cmd StartBasket) error {
	basket := domain.StartBasket(cmd.ID)

	return h.repo.Save(ctx, basket)
}
