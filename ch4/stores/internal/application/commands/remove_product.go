package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	storeRepo   domain.StoreRepository
	productRepo domain.ProductRepository
}

func NewRemoveProductHandler(storeRepo domain.StoreRepository, productRepo domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	return h.productRepo.RemoveProduct(ctx, cmd.ID)
}
