package handlers

import (
	"eda-in-golang/depot/internal/domain"
	"eda-in-golang/internal/ddd"
)

func RegisterOrderHandlers[T ddd.AggregateEvent](orderHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(orderHandlers, domain.ShoppingListCompletedEvent)
}
