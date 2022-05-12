package paymentspb

import (
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/internal/registry/serdes"
)

const (
	InvoiceAggregateChannel = "mallbots.payments.events.Invoice"

	InvoicePaidEvent = "paymentsapi.InvoicePaid"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Invoice events
	if err := serde.Register(&InvoicePaid{}); err != nil {
		return err
	}

	return nil
}

func (*InvoicePaid) Key() string { return InvoicePaidEvent }
