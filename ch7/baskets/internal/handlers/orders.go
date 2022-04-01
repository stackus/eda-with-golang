package handlers

import (
	"eda-in-golang/ch7/baskets/internal/domain"
	"eda-in-golang/ch7/internal/ddd"
)

func RegisterOrderHandlers[T ddd.AggregateEvent](orderHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(domain.BasketCheckedOutEvent, orderHandlers)
}
