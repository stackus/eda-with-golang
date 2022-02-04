package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type BuildShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type BuildShoppingListHandler struct {
	listRepo    domain.ShoppingListRepository
	storeRepo   domain.StoreRepository
	productRepo domain.ProductRepository
}

func NewSubmitOrderHandler(listRepo domain.ShoppingListRepository, storeRepo domain.StoreRepository, productRepo domain.ProductRepository) BuildShoppingListHandler {
	return BuildShoppingListHandler{
		listRepo:    listRepo,
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (h BuildShoppingListHandler) BuildShoppingList(ctx context.Context, cmd BuildShoppingList) error {
	list := domain.CreateShoppingList(cmd.ID, cmd.OrderID)

	for _, item := range cmd.Items {
		store, err := h.storeRepo.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		product, err := h.productRepo.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		err = list.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}

	return errors.Wrap(h.listRepo.Save(ctx, list), "building shopping list")
}
