package ordering

import (
	"context"

	"eda-in-golang/ch9/baskets/basketspb"
	"eda-in-golang/ch9/customers/customerspb"
	"eda-in-golang/ch9/depot/depotpb"
	"eda-in-golang/ch9/internal/ac"
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
	"eda-in-golang/ch9/payments/paymentspb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = basketspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = depotpb.Registrations(reg); err != nil {
		return err
	}
	if err = paymentspb.Registrations(reg); err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("ordering.events", mono.DB(), reg),
		// es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("ordering.snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	sagaStore := pg.NewSagaStore("ordering.sagas", mono.DB(), reg)
	createOrderSagaRepo := ac.NewSagaRepository[*domain.CreateOrderData](reg, sagaStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	// customers := grpc.NewCustomerRepository(conn)
	// payments := grpc.NewPaymentRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(orders, shopping, domainDispatcher),
		mono.Logger(),
	)
	createOrderSaga := logging.LogReplyHandlerAccess[*domain.CreateOrderData](
		ac.NewOrchestrator[*domain.CreateOrderData](application.NewCreateOrderSaga(), createOrderSagaRepo, commandStream),
		"CreateOrderSaga", mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(app, createOrderSaga),
		"IntegrationEvents", mono.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterCreateOrderReplies(replyStream, createOrderSaga); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err = serde.Register(domain.Order{}); err != nil {
		return err
	}
	// order events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderRejected{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderApproved{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err = serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	// Saga data
	if err = serde.RegisterKey(application.CreateOrderSagaName, domain.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}
