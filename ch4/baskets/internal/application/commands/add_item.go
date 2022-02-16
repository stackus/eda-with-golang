package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type AddItem struct {
	ID        domain.BasketID
	ProductID domain.ProductID
	Quantity  int
}

type AddItemHandler struct {
	baskets  domain.BasketRepository
	stores   domain.StoreRepository
	products domain.ProductRepository
}

func NewAddItemHandler(baskets domain.BasketRepository, stores domain.StoreRepository, products domain.ProductRepository) AddItemHandler {
	return AddItemHandler{
		baskets:  baskets,
		stores:   stores,
		products: products,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	basket, err := h.baskets.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	product, err := h.products.Find(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	store, err := h.stores.Find(ctx, product.StoreID)
	if err != nil {
		return nil
	}
	err = basket.AddItem(store, product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.baskets.Update(ctx, basket)
}
