package domain

import (
	"context"
)

type BasketRepository interface {
	Find(ctx context.Context, basketID BasketID) (*Basket, error)
	Save(ctx context.Context, basket *Basket) error
	Update(ctx context.Context, basket *Basket) error
}
