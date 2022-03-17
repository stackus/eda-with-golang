package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/baskets/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOutEvent, orderHandlers)
}
