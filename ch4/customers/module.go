package customers

import (
	"context"

	"eda-in-golang/ch4/customers/internal/application"
	"eda-in-golang/ch4/customers/internal/grpc"
	"eda-in-golang/ch4/customers/internal/logging"
	"eda-in-golang/ch4/customers/internal/postgres"
	"eda-in-golang/ch4/customers/internal/rest"
	"eda-in-golang/ch4/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	var app application.App
	app = application.New(customers)
	app = logging.LogApplicationAccess(app, mono.Logger())

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
