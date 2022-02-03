package domain

import (
	"context"
)

type ProductRepository interface {
	Find(ctx context.Context, storeID, productID string) (*Product, error)
}
