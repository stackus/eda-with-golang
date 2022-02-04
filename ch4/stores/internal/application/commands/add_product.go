package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type AddProduct struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type AddProductHandler struct {
	storeRepo   domain.StoreRepository
	productRepo domain.ProductRepository
}

func NewAddProductHandler(storeRepo domain.StoreRepository, productRepo domain.ProductRepository) AddProductHandler {
	return AddProductHandler{
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (h AddProductHandler) AddProduct(ctx context.Context, cmd AddProduct) error {
	_, err := h.storeRepo.Find(ctx, cmd.StoreID)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	product, err := domain.CreateProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	return errors.Wrap(h.productRepo.AddProduct(ctx, product), "error adding product")
}
