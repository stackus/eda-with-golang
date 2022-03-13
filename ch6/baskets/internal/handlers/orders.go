package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/baskets/internal/application"
	"github.com/stackus/eda-with-golang/ch6/baskets/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
