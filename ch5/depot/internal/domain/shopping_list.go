package domain

import (
	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch5/internal/ddd"
)

var (
	ErrShoppingCannotBeCanceled = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be canceled")
)

type ShoppingList struct {
	ddd.AggregateBase
	OrderID       string
	Stops         Stops
	AssignedBotID string
	Status        ShoppingListStatus
}

func CreateShopping(id, orderID string) *ShoppingList {
	shoppingList := &ShoppingList{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		OrderID: orderID,
		Status:  ShoppingListIsAvailable,
		Stops:   make(Stops),
	}

	return shoppingList
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	if _, exists := sl.Stops[store.ID]; !exists {
		sl.Stops[store.ID] = &Stop{
			StoreName:     store.Name,
			StoreLocation: store.Location,
			Items:         make(Items),
		}
	}

	return sl.Stops[store.ID].AddItem(product, quantity)
}

func (sl *ShoppingList) Cancel() error {
	// validate status

	sl.Status = ShoppingListIsCanceled

	return nil
}

func (sl *ShoppingList) Assign(id string) error {
	// validate status

	sl.AssignedBotID = id
	sl.Status = ShoppingListIsAssigned

	return nil
}

func (sl *ShoppingList) Complete() error {
	// validate status

	sl.Status = ShoppingListIsCompleted

	return nil
}
