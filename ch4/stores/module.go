package stores

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	storeRepo := postgres.NewStoreRepository("store.stores", mono.DB())
	participatingStoreRepo := postgres.NewParticipatingStoreRepository("store.stores", mono.DB())
	offeringRepo := postgres.NewOfferingRepository("store.offerings", mono.DB())

	// setup application
	app := application.New(storeRepo, participatingStoreRepo, offeringRepo)

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, app, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	return nil
}
