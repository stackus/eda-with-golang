package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type GetBasket struct {
	ID string
}

type GetBasketHandler struct {
	repo ports.BasketRepository
}

func NewGetBasketHandler(repo ports.BasketRepository) GetBasketHandler {
	return GetBasketHandler{repo: repo}
}

func (h GetBasketHandler) GetBasket(ctx context.Context, query GetBasket) (*domain.Basket, error) {
	return h.repo.FindBasket(ctx, query.ID)
}
