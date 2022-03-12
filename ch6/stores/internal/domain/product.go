package domain

import (
	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

const ProductAggregate = "stores.Product"

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	ddd.Aggregate
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func NewProduct(id string) *Product {
	return &Product{
		Aggregate: ddd.NewAggregate(id, ProductAggregate),
	}
}

func CreateProduct(id, storeID, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := NewProduct(id)
	product.StoreID = storeID
	product.Name = name
	product.Description = description
	product.SKU = sku
	product.Price = price

	product.AddEvent(ProductAddedEvent, &ProductAdded{
		Product: product,
	})

	return product, nil
}

func (p *Product) Remove() error {
	p.AddEvent(ProductRemovedEvent, &ProductRemoved{
		Product: p,
	})

	return nil
}
