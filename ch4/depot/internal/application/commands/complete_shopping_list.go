package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type CompleteShoppingList struct {
	ID domain.ShoppingListID
}

type CompleteShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
}

func NewCompleteShoppingListHandler(shoppingLists domain.ShoppingListRepository) CompleteShoppingListHandler {
	return CompleteShoppingListHandler{
		shoppingLists: shoppingLists,
	}
}

func (h CompleteShoppingListHandler) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Complete()
	if err != nil {
		return err
	}

	return h.shoppingLists.Update(ctx, list)
}
