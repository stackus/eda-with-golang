package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type BuildShoppingList struct {
	ID      string
	OrderID string
	Items   []domain.OrderItem
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
	svc := domain.NewShoppingListService(h.storeRepo, h.productRepo)

	shoppingList, err := svc.BuildShoppingList(ctx, cmd.ID, cmd.OrderID, cmd.Items)
	if err != nil {
		return err
	}

	return h.listRepo.Save(ctx, shoppingList)
}
