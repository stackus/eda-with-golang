package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/stores/internal/domain"
)

func RegisterCatalogHandlers[T ddd.AggregateEvent](catalogHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(catalogHandlers,
		domain.ProductAddedEvent,
		domain.ProductRebrandedEvent,
		domain.ProductPriceIncreasedEvent,
		domain.ProductPriceDecreasedEvent,
		domain.ProductRemovedEvent,
	)
}
