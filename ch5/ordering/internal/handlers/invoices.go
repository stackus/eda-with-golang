package handlers

import (
	"github.com/stackus/eda-with-golang/ch5/internal/ddd"
	"github.com/stackus/eda-with-golang/ch5/ordering/internal/application"
	"github.com/stackus/eda-with-golang/ch5/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
