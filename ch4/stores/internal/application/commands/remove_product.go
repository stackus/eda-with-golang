package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	stores   domain.StoreRepository
	products domain.ProductRepository
}

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{
		stores:   stores,
		products: products,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	return h.products.RemoveProduct(ctx, cmd.ID)
}
