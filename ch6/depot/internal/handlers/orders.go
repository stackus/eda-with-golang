package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/depot/internal/application"
	"github.com/stackus/eda-with-golang/ch6/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers.OnShoppingListCompleted)
}