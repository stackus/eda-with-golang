package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ProductAddedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRebrandedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceIncreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceDecreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRemovedEvent, catalogHandlers)
}
