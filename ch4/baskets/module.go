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
	basketRepo := postgres.NewBasketRepository("basket.baskets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	productRepo := grpc.NewProductRepository(conn)
	orderRepo := grpc.NewOrderRepository(conn)

	// setup application
	app := application.New(basketRepo, productRepo, orderRepo)

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, app, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	return
}
