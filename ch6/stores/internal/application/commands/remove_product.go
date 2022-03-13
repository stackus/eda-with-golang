package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	products domain.ProductRepository
}

func NewRemoveProductHandler(products domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{
		products: products,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Remove(); err != nil {
		return err
	}

	if err = h.products.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	return nil
}
