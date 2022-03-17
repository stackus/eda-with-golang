package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers)
}
