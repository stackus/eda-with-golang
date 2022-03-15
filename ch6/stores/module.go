package stores

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	es2 "github.com/stackus/eda-with-golang/ch6/internal/es"
	"github.com/stackus/eda-with-golang/ch6/internal/monolith"
	pg "github.com/stackus/eda-with-golang/ch6/internal/postgres"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
	"github.com/stackus/eda-with-golang/ch6/internal/registry/serdes"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/es"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/grpc"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/handlers"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/logging"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/postgres"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher()
	aggregateStore := es2.NewEventPublisher(
		pg.NewSnapshotStore(
			pg.NewEventStore("stores.events", mono.DB(), reg),
			"stores.snapshots", mono.DB(), reg,
		),
		domainDispatcher,
	)
	stores := es.NewStoreRepository(reg, aggregateStore)
	products := es.NewProductRepository(reg, aggregateStore)
	catalog := postgres.NewCatalogRepository("stores.products", mono.DB())
	mall := postgres.NewMallRepository("stores.stores", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(stores, products, catalog, mall),
		mono.Logger(),
	)
	catalogHandlers := logging.LogDomainEventHandlerAccess(
		application.NewCatalogHandlers(catalog),
		"Catalog", mono.Logger(),
	)
	mallHandlers := logging.LogDomainEventHandlerAccess(
		application.NewMallHandlers(mall),
		"Mall", mono.Logger(),
	)

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
	handlers.RegisterCatalogHandlers(catalogHandlers, domainDispatcher)
	handlers.RegisterMallHandlers(mallHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(domain.Store{}); err != nil {
		return
	}
	// store events
	if err = serde.Register(domain.StoreCreated{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.StoreParticipationEnabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.StoreParticipationDisabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.Register(domain.StoreRebranded{}); err != nil {
		return
	}
	// store snapshots
	if err = serde.RegisterKey(domain.StoreV1{}.SnapshotName(), domain.StoreV1{}); err != nil {
		return
	}

	// Product
	if err = serde.Register(domain.Product{}); err != nil {
		return
	}
	// product events
	if err = serde.Register(domain.ProductAdded{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRebranded{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.ProductPriceIncreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.ProductPriceDecreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRemoved{}); err != nil {
		return
	}
	// product snapshots
	if err = serde.RegisterKey(domain.ProductV1{}.SnapshotName(), domain.ProductV1{}); err != nil {
		return
	}

	return
}
