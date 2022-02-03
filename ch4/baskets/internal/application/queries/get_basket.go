package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type GetBasket struct {
	ID string
}

type GetBasketHandler struct {
	repo domain.BasketRepository
}

func NewGetBasketHandler(repo domain.BasketRepository) GetBasketHandler {
	return GetBasketHandler{repo: repo}
}

func (h GetBasketHandler) GetBasket(ctx context.Context, query GetBasket) (*domain.Basket, error) {
	return h.repo.Find(ctx, query.ID)
}
