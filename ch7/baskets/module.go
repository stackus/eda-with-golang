package baskets

import (
	"context"

	"eda-in-golang/ch7/baskets/internal/application"
	"eda-in-golang/ch7/baskets/internal/grpc"
	"eda-in-golang/ch7/baskets/internal/handlers"
	"eda-in-golang/ch7/baskets/internal/logging"
	"eda-in-golang/ch7/baskets/internal/postgres"
	"eda-in-golang/ch7/baskets/internal/rest"
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	baskets := postgres.NewBasketRepository("baskets.baskets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, orders, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess(
		application.NewOrderHandlers(orders),
		"Order", mono.Logger(),
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
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return
}
