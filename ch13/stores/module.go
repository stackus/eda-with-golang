package stores

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/es"
	"eda-in-golang/internal/jetstream"
	pg "eda-in-golang/internal/postgres"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/registry/serdes"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/tm"
	"eda-in-golang/stores/internal/application"
	"eda-in-golang/stores/internal/domain"
	"eda-in-golang/stores/internal/grpc"
	"eda-in-golang/stores/internal/handlers"
	"eda-in-golang/stores/internal/postgres"
	"eda-in-golang/stores/internal/rest"
	"eda-in-golang/stores/storespb"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return svc.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), c.Get("logger").(zerolog.Logger)), nil
	})
	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return svc.DB(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.MessageStream),
			pg.NewOutboxStore("stores.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("messagePublisher", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("stores.outbox", tx)
		return am.NewMessagePublisher(
			c.Get("stream").(am.MessageStream),
			am.OtelMessageContextInjector(),
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	container.AddSingleton("messageSubscriber", func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			c.Get("stream").(am.MessageStream),
			am.OtelMessageContextExtractor(),
		), nil
	})
	container.AddScoped("eventPublisher", func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get("registry").(registry.Registry),
			c.Get("messagePublisher").(am.MessagePublisher),
		), nil
	})
	container.AddScoped("inboxStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		return pg.NewInboxStore("stores.inbox", tx), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			pg.NewEventStore("stores.events", tx, reg),
			pg.NewSnapshotStore("stores.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Store](
			domain.StoreAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Product](
			domain.ProductAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})
	container.AddScoped("catalog", func(c di.Container) (any, error) {
		return postgres.NewCatalogRepository("stores.products", c.Get("tx").(*sql.Tx)), nil
	})
	container.AddScoped("mall", func(c di.Container) (any, error) {
		return postgres.NewMallRepository("stores.stores", c.Get("tx").(*sql.Tx)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return application.New(
			c.Get("stores").(domain.StoreRepository),
			c.Get("products").(domain.ProductRepository),
			c.Get("catalog").(domain.CatalogRepository),
			c.Get("mall").(domain.MallRepository),
			c.Get("domainDispatcher").(ddd.EventPublisher[ddd.Event]),
		), nil
	})
	container.AddScoped("catalogHandlers", func(c di.Container) (any, error) {
		return handlers.NewCatalogHandlers(c.Get("catalog").(domain.CatalogRepository)), nil
	})
	container.AddScoped("mallHandlers", func(c di.Container) (any, error) {
		return handlers.NewMallHandlers(c.Get("mall").(domain.MallRepository)), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get("eventPublisher").(am.EventPublisher)), nil
	})

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, svc.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, svc.Mux(), svc.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(svc.Mux()); err != nil {
		return err
	}
	handlers.RegisterCatalogHandlersTx(container)
	handlers.RegisterMallHandlersTx(container)
	handlers.RegisterDomainEventHandlersTx(container)
	if err = storespb.RegisterAsyncAPI(svc.Mux()); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(domain.Store{}, func(v any) error {
		store := v.(*domain.Store)
		store.Aggregate = es.NewAggregate("", domain.StoreAggregate)
		return nil
	}); err != nil {
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
	if err = serde.Register(domain.Product{}, func(v any) error {
		store := v.(*domain.Product)
		store.Aggregate = es.NewAggregate("", domain.ProductAggregate)
		return nil
	}); err != nil {
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
func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("stores outbox processor encountered an error")
		}
	}()
}
