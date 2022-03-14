package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers application.DomainEventHandlers,
	domainSubscriber ddd.EventSubscriber,
) {
	domainSubscriber.Subscribe(domain.ProductAddedEvent, catalogHandlers.OnProductAdded)
	domainSubscriber.Subscribe(domain.ProductRebrandedEvent, catalogHandlers.OnProductRebranded)
	domainSubscriber.Subscribe(domain.ProductPriceIncreasedEvent, catalogHandlers.OnProductPriceIncreased)
	domainSubscriber.Subscribe(domain.ProductPriceDecreasedEvent, catalogHandlers.OnProductPriceDecreased)
	domainSubscriber.Subscribe(domain.ProductRemovedEvent, catalogHandlers.OnProductRemoved)
}
