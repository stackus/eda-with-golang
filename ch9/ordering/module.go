package ordering

import (
	"context"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/internal/es"
	"eda-in-golang/ch9/internal/jetstream"
	"eda-in-golang/ch9/internal/monolith"
	pg "eda-in-golang/ch9/internal/postgres"
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/internal/registry/serdes"
	"eda-in-golang/ch9/ordering/internal/application"
	"eda-in-golang/ch9/ordering/internal/domain"
	"eda-in-golang/ch9/ordering/internal/grpc"
	"eda-in-golang/ch9/ordering/internal/handlers"
	"eda-in-golang/ch9/ordering/internal/logging"
	"eda-in-golang/ch9/ordering/internal/rest"
	"eda-in-golang/ch9/ordering/orderingpb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("ordering.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("ordering.snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, shopping)
	app = logging.LogApplicationAccess(app, mono.Logger())
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
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

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err := serde.Register(domain.Order{}); err != nil {
		return err
	}
	// order events
	if err := serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err := serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}
