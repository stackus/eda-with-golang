package domain

import (
	"fmt"
)

var (
	ErrShoppingCannotBeCancelled = fmt.Errorf("the shopping list cannot be cancelled")
)

type ShoppingListID string

type ShoppingListStatus string

const (
	ShoppingListUnknown   ShoppingListStatus = ""
	ShoppingListAvailable ShoppingListStatus = "available"
	ShoppingListAssigned  ShoppingListStatus = "assigned"
	ShoppingListActive    ShoppingListStatus = "active"
	ShoppingListCompleted ShoppingListStatus = "completed"
	ShoppingListCancelled ShoppingListStatus = "cancelled"
)

func (i ShoppingListID) String() string {
	return string(i)
}

func ToShoppingListID(id string) ShoppingListID {
	return ShoppingListID(id)
}

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListAvailable, ShoppingListAssigned, ShoppingListActive, ShoppingListCompleted, ShoppingListCancelled:
		return string(s)
	default:
		return ""
	}
}

func ToShoppingListStatus(status string) ShoppingListStatus {
	switch status {
	case ShoppingListAvailable.String():
		return ShoppingListAvailable
	case ShoppingListAssigned.String():
		return ShoppingListAssigned
	case ShoppingListActive.String():
		return ShoppingListActive
	case ShoppingListCompleted.String():
		return ShoppingListCompleted
	case ShoppingListCancelled.String():
		return ShoppingListCancelled
	default:
		return ShoppingListUnknown
	}
}

type ShoppingList struct {
	ID            ShoppingListID
	OrderID       string
	Stops         *Stops
	AssignedBotID BotID
	Status        ShoppingListStatus
}

func CreateShopping(id ShoppingListID, orderID string) *ShoppingList {
	return &ShoppingList{
		ID:      id,
		OrderID: orderID,
		Status:  ShoppingListAvailable,
	}
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	return sl.Stops.AddItem(store, product, quantity)
}

func (sl *ShoppingList) Cancel() error {
	// validate status

	sl.Status = ShoppingListCancelled

	return nil
}

func (sl *ShoppingList) Assign(id BotID) error {
	// validate status

	sl.AssignedBotID = id
	sl.Status = ShoppingListAssigned

	return nil
}

func (sl *ShoppingList) Complete() error {
	// validate status

	sl.Status = ShoppingListCompleted

	return nil
}
