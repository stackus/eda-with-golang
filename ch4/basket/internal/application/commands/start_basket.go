package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type StartBasket struct {
	ID string
}

type StartBasketHandler struct {
	repo ports.BasketRepository
}

func NewStartBasketHandler(repo ports.BasketRepository) StartBasketHandler {
	return StartBasketHandler{repo: repo}
}

func (h StartBasketHandler) StartBasket(ctx context.Context, cmd StartBasket) error {
	basket := domain.StartBasket(cmd.ID)

	return h.repo.Save(ctx, basket)
}
