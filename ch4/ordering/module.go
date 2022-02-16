package ordering

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	invoices := grpc.NewInvoiceRepository(conn)
	shoppingLists := grpc.NewShoppingListRepository(conn)

	// setup application
	app := application.New(orders, invoices, shoppingLists)

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

	return nil
}
