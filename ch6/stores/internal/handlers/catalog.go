package handlers

import (
	"eda-in-golang/ch6/internal/ddd"
	"eda-in-golang/ch6/stores/internal/domain"
)

func RegisterCatalogHandlers[T ddd.AggregateEvent](catalogHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(domain.ProductAddedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRebrandedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceIncreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceDecreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRemovedEvent, catalogHandlers)
}
