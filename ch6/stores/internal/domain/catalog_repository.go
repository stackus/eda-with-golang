package domain

import (
	"context"
)

type CatalogRepository interface {
	AddProduct(ctx context.Context, productID, storeID, name, description, sku string, price float64) error
	RemoveProduct(ctx context.Context, productID string) error
	Find(ctx context.Context, productID string) (*Product, error)
	GetCatalog(ctx context.Context, storeID string) ([]*Product, error)
}
