package stores

import (
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(mono monolith.Monolith) error {
	// Startup Driven adapters
	storeRepo := postgres.NewStoreRepository("store.stores", mono.DB())

	// Startup application
	app := application.New(storeRepo)

	// Setup Driver adapters
	if err := grpc.Register(app, mono); err != nil {
		return err
	}
	if err := rest.Register(app, mono); err != nil {
		return err
	}

	return nil
}
