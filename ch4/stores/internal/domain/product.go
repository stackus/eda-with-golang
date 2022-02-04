package domain

import (
	"fmt"
)

var (
	ErrProductNameIsBlank     = fmt.Errorf("the product name cannot be blank")
	ErrProductPriceIsNegative = fmt.Errorf("the product price cannot be negative")
)

type Product struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(id, storeID, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := &Product{
		ID:          id,
		StoreID:     storeID,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	}

	return product, nil
}
