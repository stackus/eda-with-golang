package domain

import (
	"context"
)

type ShoppingListRepository interface {
	Find(ctx context.Context, id ShoppingListID) (*ShoppingList, error)
	FindByOrderID(ctx context.Context, orderID string) (*ShoppingList, error)
	Save(ctx context.Context, list *ShoppingList) error
	Update(ctx context.Context, list *ShoppingList) error
}
