package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type BasketRepository interface {
	Find(ctx context.Context, basketID string) (*domain.Basket, error)
	Save(ctx context.Context, basket *domain.Basket) error
	Update(ctx context.Context, basket *domain.Basket) error
}
