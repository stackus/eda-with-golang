package payments

import (
	"context"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/internal/jetstream"
	"eda-in-golang/ch9/internal/monolith"
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/ordering/orderingpb"
	"eda-in-golang/ch9/payments/internal/application"
	"eda-in-golang/ch9/payments/internal/grpc"
	"eda-in-golang/ch9/payments/internal/handlers"
	"eda-in-golang/ch9/payments/internal/logging"
	"eda-in-golang/ch9/payments/internal/postgres"
	"eda-in-golang/ch9/payments/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	invoices := postgres.NewInvoiceRepository("payments.invoices", mono.DB())
	payments := postgres.NewPaymentRepository("payments.payments", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(invoices, payments, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewOrderHandlers(app),
		"Order", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterOrderHandlers(orderHandlers, eventStream); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers[ddd.Event](integrationEventHandlers, domainDispatcher)

	return nil
}
