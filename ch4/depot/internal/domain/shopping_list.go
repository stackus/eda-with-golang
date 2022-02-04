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

func CreateShoppingList(id, orderID string) *ShoppingList {
	return &ShoppingList{
		ID:      id,
		OrderID: orderID,
		Status:  ShoppingListAvailable,
	}
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	for _, stop := range sl.Stops {
		if stop.StoreID == store.ID {
			stop.Items = append(stop.Items, &Item{
				ID:       product.ID,
				Name:     product.Name,
				Quantity: quantity,
			})
			return nil
		}
	}
	sl.Stops = append(sl.Stops, &Stop{
		StoreID:       store.ID,
		StoreName:     store.ID,
		StoreLocation: store.Location,
		Items: []*Item{
			{
				ID:       product.ID,
				Name:     product.Name,
				Quantity: quantity,
			},
		},
	})

	return nil
}

func (sl *ShoppingList) Cancel() error {
	sl.Status = ShoppingListCancelled

	return nil
}
