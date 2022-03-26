package customers

import (
	"context"

	"eda-in-golang/ch7/customers/internal/application"
	"eda-in-golang/ch7/customers/internal/grpc"
	"eda-in-golang/ch7/customers/internal/logging"
	"eda-in-golang/ch7/customers/internal/postgres"
	"eda-in-golang/ch7/customers/internal/rest"
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers, domainDispatcher),
		mono.Logger(),
	)

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
