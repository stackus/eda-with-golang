package handlers

import (
	"eda-in-golang/ch5/baskets/internal/application"
	"eda-in-golang/ch5/baskets/internal/domain"
	"eda-in-golang/ch5/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
