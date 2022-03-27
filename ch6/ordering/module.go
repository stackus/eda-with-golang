package ordering

import (
	"context"

	"eda-in-golang/ch6/internal/ddd"
	es2 "eda-in-golang/ch6/internal/es"
	"eda-in-golang/ch6/internal/monolith"
	pg "eda-in-golang/ch6/internal/postgres"
	"eda-in-golang/ch6/internal/registry"
	"eda-in-golang/ch6/internal/registry/serdes"
	"eda-in-golang/ch6/ordering/internal/application"
	"eda-in-golang/ch6/ordering/internal/domain"
	"eda-in-golang/ch6/ordering/internal/es"
	"eda-in-golang/ch6/ordering/internal/grpc"
	"eda-in-golang/ch6/ordering/internal/handlers"
	"eda-in-golang/ch6/ordering/internal/logging"
	"eda-in-golang/ch6/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	err = registrations(reg)
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher()
	aggregateStore := es2.AggregateStoreWithMiddleware(
		pg.NewEventStore("baskets.events", mono.DB(), reg),
		es2.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("baskets.snapshots", mono.DB(), reg),
	)
	orders := es.NewOrderRepository(reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, shopping)
	app = logging.LogApplicationAccess(app, mono.Logger())
	// setup application handlers
	notificationHandlers := logging.LogEventHandlerAccess(
		application.NewNotificationHandlers(notifications),
		"Notification", mono.Logger(),
	)
	invoiceHandlers := logging.LogEventHandlerAccess(
		application.NewInvoiceHandlers(invoices),
		"Invoice", mono.Logger(),
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
	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)
	handlers.RegisterInvoiceHandlers(invoiceHandlers, domainDispatcher)

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
