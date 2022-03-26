package ordering

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/monolith"
	"eda-in-golang/ch7/ordering/internal/application"
	"eda-in-golang/ch7/ordering/internal/grpc"
	"eda-in-golang/ch7/ordering/internal/handlers"
	"eda-in-golang/ch7/ordering/internal/logging"
	"eda-in-golang/ch7/ordering/internal/postgres"
	"eda-in-golang/ch7/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
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
	app = application.New(orders, customers, payments, shopping, domainDispatcher)
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
