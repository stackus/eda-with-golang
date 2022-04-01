package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/ordering/internal/domain"
)

func RegisterInvoiceHandlers[T ddd.AggregateEvent](invoiceHandlers ddd.EventHandler[T], domainSubscriber ddd.EventSubscriber[T]) {
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, invoiceHandlers)
}
