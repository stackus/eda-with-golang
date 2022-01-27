package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type BasketRepository interface {
	FindBasket(ctx context.Context, basketID string) (*domain.Basket, error)
	SaveBasket(ctx context.Context, basket *domain.Basket) error
	UpdateBasket(ctx context.Context, basket *domain.Basket) error
}
