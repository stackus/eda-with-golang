package domain

import (
	"context"
)

type ProductRepository interface {
	Find(ctx context.Context, id string) (*Product, error)
	Save(ctx context.Context, product *Product) error
}
