package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type RemoveItem struct {
	ID        string
	StoreID   string
	ProductID string
	Quantity  int
}

type RemoveItemHandler struct {
	basketRepo  domain.BasketRepository
	productRepo domain.ProductRepository
}

func NewRemoveItemHandler(basketRepo domain.BasketRepository, productRepo domain.ProductRepository) RemoveItemHandler {
	return RemoveItemHandler{
		basketRepo:  basketRepo,
		productRepo: productRepo,
	}
}

func (h RemoveItemHandler) RemoveItem(ctx context.Context, cmd RemoveItem) error {
	product, err := h.productRepo.Find(ctx, cmd.ProductID, cmd.StoreID)
	if err != nil {
		return err
	}

	basket, err := h.basketRepo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.RemoveItem(product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.basketRepo.Update(ctx, basket)
}
