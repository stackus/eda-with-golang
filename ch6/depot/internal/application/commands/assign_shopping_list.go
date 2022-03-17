package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type AssignShoppingList struct {
	ID    string
	BotID string
}

type AssignShoppingListHandler struct {
	shoppingLists   domain.ShoppingListRepository
	domainPublisher ddd.EventPublisher
}

func NewAssignShoppingListHandler(shoppingList domain.ShoppingListRepository, domainPublisher ddd.EventPublisher,
) AssignShoppingListHandler {
	return AssignShoppingListHandler{
		shoppingLists:   shoppingList,
		domainPublisher: domainPublisher,
	}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Assign(cmd.BotID)
	if err != nil {
		return err
	}

	if err = h.shoppingLists.Update(ctx, list); err != nil {
		return err
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
