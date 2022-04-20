package handlers

import (
	"eda-in-golang/ch8/baskets/internal/domain"
	"eda-in-golang/ch8/internal/ddd"
)

func RegisterOrderHandlers[T ddd.AggregateEvent](orderHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(orderHandlers, domain.BasketCheckedOutEvent)
}
