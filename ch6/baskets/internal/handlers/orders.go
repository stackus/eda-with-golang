package handlers

import (
	"eda-in-golang/ch6/baskets/internal/domain"
	"eda-in-golang/ch6/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOutEvent, orderHandlers)
}
