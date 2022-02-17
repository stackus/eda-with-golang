package domain

import (
	"fmt"
)

type Items map[string]*Item

type Item struct {
	ProductName string
	Quantity    int
}

func (i *Items) AddItem(product *Product, quantity int) error {
	if _, exists := (*i)[product.ID]; exists {
		return fmt.Errorf("product already added: %s", product.Name)
	}

	(*i)[product.ID] = &Item{
		ProductName: product.Name,
		Quantity:    quantity,
	}

	return nil
}
