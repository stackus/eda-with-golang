package domain

import (
	"context"
)

type ShoppingListRepository interface {
	Save(ctx context.Context, order *Order) error
	Delete(ctx context.Context, orderID OrderID) error
}
