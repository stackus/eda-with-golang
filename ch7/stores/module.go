package stores

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/em"
	"eda-in-golang/ch7/internal/es"
	"eda-in-golang/ch7/internal/jetstream"
	"eda-in-golang/ch7/internal/monolith"
	pg "eda-in-golang/ch7/internal/postgres"
	"eda-in-golang/ch7/internal/registry"
	"eda-in-golang/ch7/internal/registry/serdes"
	"eda-in-golang/ch7/stores/internal/application"
	"eda-in-golang/ch7/stores/internal/domain"
	"eda-in-golang/ch7/stores/internal/grpc"
	"eda-in-golang/ch7/stores/internal/handlers"
	"eda-in-golang/ch7/stores/internal/logging"
	"eda-in-golang/ch7/stores/internal/postgres"
	"eda-in-golang/ch7/stores/internal/rest"
	"eda-in-golang/ch7/stores/storespb"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := em.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("stores.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("stores.snapshots", mono.DB(), reg),
	)
	stores := es.NewAggregateRepository[*domain.Store](domain.StoreAggregate, reg, aggregateStore)
	products := es.NewAggregateRepository[*domain.Product](domain.ProductAggregate, reg, aggregateStore)
	catalog := postgres.NewCatalogRepository("stores.products", mono.DB())
	mall := postgres.NewMallRepository("stores.stores", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(stores, products, catalog, mall),
		mono.Logger(),
	)
	catalogHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewCatalogHandlers(catalog),
		"Catalog", mono.Logger(),
	)
	mallHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewMallHandlers(mall),
		"Mall", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
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
	handlers.RegisterCatalogHandlers[ddd.AggregateEvent](catalogHandlers, domainDispatcher)
	handlers.RegisterMallHandlers[ddd.AggregateEvent](mallHandlers, domainDispatcher)
	handlers.RegisterIntegrationEventHandlers[ddd.AggregateEvent](integrationEventHandlers, domainDispatcher)

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
