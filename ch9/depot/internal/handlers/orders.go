package handlers

import (
	"eda-in-golang/ch9/depot/internal/domain"
	"eda-in-golang/ch9/internal/ddd"
)

func RegisterOrderHandlers[T ddd.AggregateEvent](orderHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(orderHandlers, domain.ShoppingListCompletedEvent)
}
