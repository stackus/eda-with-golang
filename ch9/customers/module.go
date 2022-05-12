package customers

import (
	"context"

	"eda-in-golang/ch9/customers/customerspb"
	"eda-in-golang/ch9/customers/internal/application"
	"eda-in-golang/ch9/customers/internal/grpc"
	"eda-in-golang/ch9/customers/internal/handlers"
	"eda-in-golang/ch9/customers/internal/logging"
	"eda-in-golang/ch9/customers/internal/postgres"
	"eda-in-golang/ch9/customers/internal/rest"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/internal/jetstream"
	"eda-in-golang/ch9/internal/monolith"
	"eda-in-golang/ch9/internal/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers, domainDispatcher),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers[ddd.AggregateEvent](integrationEventHandlers, domainDispatcher)

	return nil
}
