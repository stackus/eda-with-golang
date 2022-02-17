package baskets

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/rest"
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	baskets := postgres.NewBasketRepository("basket.baskets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := application.New(baskets, stores, products, orders)

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

	return
}
