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
	// get product data from the store module
	product, err := h.productRepo.FindProduct(ctx, cmd.ProductID, cmd.StoreID)
	if err != nil {
		return err
	}

	basket, err := h.basketRepo.FindBasket(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basket.AddItem(product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.basketRepo.UpdateBasket(ctx, basket)
}
