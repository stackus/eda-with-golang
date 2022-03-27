package handlers

import (
	"eda-in-golang/ch6/internal/ddd"
	"eda-in-golang/ch6/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, invoiceHandlers)
}
