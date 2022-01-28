package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type ProductRepository interface {
	Find(ctx context.Context, storeID, productID string) (*domain.Product, error)
}
