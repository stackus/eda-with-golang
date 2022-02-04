package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type AddItem struct {
	ID        string
	ProductID string
	Quantity  int
}

type AddItemHandler struct {
	basketRepo  domain.BasketRepository
	productRepo domain.ProductRepository
}

func NewAddItemHandler(basketRepo domain.BasketRepository, productRepo domain.ProductRepository) AddItemHandler {
	return AddItemHandler{
		basketRepo:  basketRepo,
		productRepo: productRepo,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	product, err := h.productRepo.Find(ctx, cmd.ProductID)
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
