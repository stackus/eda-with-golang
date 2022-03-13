package domain

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/es"
)

const ProductAggregate = "stores.Product"

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	es.Aggregate
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func NewProduct(id string) *Product {
	return &Product{
		Aggregate: es.NewAggregate(id, ProductAggregate),
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

	product.AddEvent(&ProductAdded{
		StoreID:     storeID,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	})

	return product, nil
}

func (p *Product) Remove() error {
	p.AddEvent(&ProductRemoved{})

	return nil
}

func (p *Product) ApplyEvent(ctx context.Context, event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *ProductAdded:
		p.StoreID = payload.StoreID
		p.Name = payload.Name
		p.Description = payload.Description
		p.SKU = payload.SKU
		p.Price = payload.Price

	case *ProductRemoved:
		// noop

	default:
		return errors.ErrInternal.Msgf("%T received the expected event payload %T", p, event.Payload())
	}

	return nil
}
