package handlers

import (
	"eda-in-golang/ch7/depot/internal/domain"
	"eda-in-golang/ch7/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers)
}
