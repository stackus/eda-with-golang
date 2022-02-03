package domain

import (
	"fmt"
)

var (
	ErrShoppingCannotBeCancelled = fmt.Errorf("the shopping list cannot be cancelled")
)

type ShoppingListStatus string

const (
	ShoppingListUnknown   ShoppingListStatus = ""
	ShoppingListAvailable ShoppingListStatus = "available"
	ShoppingListAssigned  ShoppingListStatus = "assigned"
	ShoppingListActive    ShoppingListStatus = "active"
	ShoppingListCompleted ShoppingListStatus = "completed"
	ShoppingListCancelled ShoppingListStatus = "cancelled"
)

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListAvailable, ShoppingListAssigned, ShoppingListActive, ShoppingListCompleted, ShoppingListCancelled:
		return string(s)
	default:
		return ""
	}
}

type ShoppingList struct {
	ID            string
	OrderID       string
	Stops         []*Stop
	AssignedBotID string
	Status        ShoppingListStatus
}

func (sl *ShoppingList) Cancel() error {
	sl.Status = ShoppingListCancelled

	return nil
}
