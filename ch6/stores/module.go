package stores

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/monolith"
	pg "github.com/stackus/eda-with-golang/ch6/internal/postgres"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
	"github.com/stackus/eda-with-golang/ch6/internal/registry/codecs"
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
	err := registerEntities(reg)
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher()
	aggregateStore := es.NewEventPublisher(
		pg.NewEventStore("stores.events", mono.DB(), reg),
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

func registerEntities(reg registry.Registry) (err error) {
	js := codecs.NewJSONCodec(reg)

	if err = js.Register(domain.StoreAggregate, domain.Store{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.StoreCreated{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.StoreParticipationEnabled{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.StoreParticipationDisabled{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.StoreRebranded{}); err != nil {
		return
	}

	if err = js.Register(domain.ProductAggregate, domain.Product{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.ProductAdded{}); err != nil {
		return
	}
	if err = ddd.RegisterEventPayload(js, domain.ProductRemoved{}); err != nil {
		return
	}

	return
}
