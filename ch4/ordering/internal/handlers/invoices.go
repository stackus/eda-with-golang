package handlers

import (
	"eda-in-golang/ch4/internal/ddd"
	"eda-in-golang/ch4/ordering/internal/application"
	"eda-in-golang/ch4/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
