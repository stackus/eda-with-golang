package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/depot/internal/domain"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
}

func NewCancelShoppingListHandler(shoppingLists domain.ShoppingListRepository) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingLists: shoppingLists,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Cancel()
	if err != nil {
		return err
	}

	return h.shoppingLists.Update(ctx, list)
}
