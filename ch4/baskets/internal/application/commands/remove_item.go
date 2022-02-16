package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type RemoveItem struct {
	ID        domain.BasketID
	ProductID domain.ProductID
	Quantity  int
}

type RemoveItemHandler struct {
	baskets  domain.BasketRepository
	products domain.ProductRepository
}

func NewRemoveItemHandler(baskets domain.BasketRepository, products domain.ProductRepository) RemoveItemHandler {
	return RemoveItemHandler{
		baskets:  baskets,
		products: products,
	}
}

func (h RemoveItemHandler) RemoveItem(ctx context.Context, cmd RemoveItem) error {
	product, err := h.products.Find(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	basket, err := h.baskets.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.RemoveItem(product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.baskets.Update(ctx, basket)
}
