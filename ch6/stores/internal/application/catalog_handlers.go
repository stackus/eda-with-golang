package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type CatalogHandlers struct {
	catalog domain.CatalogRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*CatalogHandlers)(nil)

func NewCatalogHandlers(catalog domain.CatalogRepository) *CatalogHandlers {
	return &CatalogHandlers{
		catalog: catalog,
	}
}

func (h CatalogHandlers) OnProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ProductAdded)
	return h.catalog.AddProduct(ctx, event.AggregateID(), payload.StoreID, payload.Name, payload.Description, payload.SKU, payload.Price)
}

func (h CatalogHandlers) OnProductRemoved(ctx context.Context, event ddd.Event) error {
	return h.catalog.RemoveProduct(ctx, event.AggregateID())
}
