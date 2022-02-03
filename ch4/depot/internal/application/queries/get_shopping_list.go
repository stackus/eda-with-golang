package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type GetShoppingList struct {
	ID string
}

type GetShoppingListHandler struct {
	repo domain.ShoppingListRepository
}

func NewGetShoppingListHandler(repo domain.ShoppingListRepository) GetShoppingListHandler {
	return GetShoppingListHandler{repo: repo}
}

func (h GetShoppingListHandler) GetShoppingList(ctx context.Context, query GetShoppingList) (*domain.ShoppingList, error) {
	return nil, nil
}
