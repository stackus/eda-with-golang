package application

import (
	"context"

	"eda-in-golang/ch6/internal/ddd"
	"eda-in-golang/ch6/ordering/internal/domain"
)

type InvoiceHandlers struct {
	invoices domain.InvoiceRepository
}

var _ ddd.EventHandler = (*InvoiceHandlers)(nil)

func NewInvoiceHandlers(invoices domain.InvoiceRepository) *InvoiceHandlers {
	return &InvoiceHandlers{
		invoices: invoices,
	}
}

func (h InvoiceHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	}
	return nil
}

func (h InvoiceHandlers) onOrderReadied(ctx context.Context, event ddd.Event) error {
	orderReadied := event.Payload().(*domain.OrderReadied)
	return h.invoices.Save(ctx, event.AggregateID(), orderReadied.PaymentID, orderReadied.Total)
}
