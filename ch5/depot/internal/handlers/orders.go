package handlers

import (
	"eda-in-golang/ch5/depot/internal/application"
	"eda-in-golang/ch5/depot/internal/domain"
	"eda-in-golang/ch5/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
