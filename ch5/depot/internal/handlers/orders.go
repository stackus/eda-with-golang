package handlers

import (
	"github.com/stackus/eda-with-golang/ch5/depot/internal/application"
	"github.com/stackus/eda-with-golang/ch5/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch5/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
