package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
)

type AddItem struct {
	ID        string
	StoreID   string
	ProductID string
	Quantity  int
}

type AddItemHandler struct {
	basketRepo  ports.BasketRepository
	productRepo ports.ProductRepository
}

func NewAddItemHandler(basketRepo ports.BasketRepository, productRepo ports.ProductRepository) AddItemHandler {
	return AddItemHandler{
		basketRepo:  basketRepo,
		productRepo: productRepo,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	product, err := h.productRepo.Find(ctx, cmd.StoreID, cmd.ProductID)
	if err != nil {
		return err
	}

	basket, err := h.basketRepo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.AddItem(product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.basketRepo.Update(ctx, basket)
}
