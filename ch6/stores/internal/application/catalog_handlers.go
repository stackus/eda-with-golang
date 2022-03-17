package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type CatalogHandlers struct {
	catalog domain.CatalogRepository
}

var _ ddd.EventHandler = (*CatalogHandlers)(nil)

func NewCatalogHandlers(catalog domain.CatalogRepository) *CatalogHandlers {
	return &CatalogHandlers{
		catalog: catalog,
	}
}

func (h CatalogHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case domain.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case domain.ProductPriceIncreasedEvent:
		return h.onProductPriceIncreased(ctx, event)
	case domain.ProductPriceDecreasedEvent:
		return h.onProductPriceDecreased(ctx, event)
	case domain.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h CatalogHandlers) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ProductAdded)
	return h.catalog.AddProduct(ctx, event.AggregateID(), payload.StoreID, payload.Name, payload.Description, payload.SKU, payload.Price)
}

func (h CatalogHandlers) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ProductRebranded)
	return h.catalog.Rebrand(ctx, event.AggregateID(), payload.Name, payload.Description)
}

func (h CatalogHandlers) onProductPriceIncreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.catalog.UpdatePrice(ctx, event.AggregateID(), payload.Price)
}

func (h CatalogHandlers) onProductPriceDecreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.catalog.UpdatePrice(ctx, event.AggregateID(), payload.Price)
}

func (h CatalogHandlers) onProductRemoved(ctx context.Context, event ddd.Event) error {
	return h.catalog.RemoveProduct(ctx, event.AggregateID())
}
