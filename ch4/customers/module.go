package customers

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/customers/internal/application"
	"github.com/stackus/eda-with-golang/ch4/customers/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/customers/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/customers/internal/rest"
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	app := application.New(customers)

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
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
