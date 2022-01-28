package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
)

type RemoveItem struct {
	ID        string
	StoreID   string
	ProductID string
	Quantity  int
}

type RemoveItemHandler struct {
	basketRepo  ports.BasketRepository
	productRepo ports.ProductRepository
}

func NewRemoveItemHandler(basketRepo ports.BasketRepository, productRepo ports.ProductRepository) RemoveItemHandler {
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
