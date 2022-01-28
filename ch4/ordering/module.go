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
	orderRepo := postgres.NewOrderRepository("ordering.orders", mono.DB())

	app := application.New(orderRepo)

	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, app, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	return nil
}
