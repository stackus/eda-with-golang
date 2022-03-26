package stores

import (
	"context"

	"eda-in-golang/ch5/internal/ddd"
	"eda-in-golang/ch5/internal/monolith"
	"eda-in-golang/ch5/stores/internal/application"
	"eda-in-golang/ch5/stores/internal/grpc"
	"eda-in-golang/ch5/stores/internal/logging"
	"eda-in-golang/ch5/stores/internal/postgres"
	"eda-in-golang/ch5/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	stores := postgres.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := postgres.NewProductRepository("stores.products", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

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
