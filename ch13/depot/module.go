package depot

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/depot/depotpb"
	"eda-in-golang/depot/internal/application"
	"eda-in-golang/depot/internal/domain"
	"eda-in-golang/depot/internal/grpc"
	"eda-in-golang/depot/internal/handlers"
	"eda-in-golang/depot/internal/postgres"
	"eda-in-golang/depot/internal/rest"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/jetstream"
	pg "eda-in-golang/internal/postgres"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/tm"
	"eda-in-golang/stores/storespb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()

	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotpb.Registrations(reg); err != nil {
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
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return svc.DB(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.MessageStream),
			pg.NewOutboxStore("depot.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("messagePublisher", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("baskets.outbox", tx)
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
	container.AddScoped("commandPublisher", func(c di.Container) (any, error) {
		return am.NewCommandPublisher(
			c.Get("registry").(registry.Registry),
			c.Get("messagePublisher").(am.MessagePublisher),
		), nil
	})
	container.AddScoped("replyPublisher", func(c di.Container) (any, error) {
		return am.NewReplyPublisher(
			c.Get("registry").(registry.Registry),
			c.Get("messagePublisher").(am.MessagePublisher),
		), nil
	})
	container.AddScoped("inboxStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		return pg.NewInboxStore("baskets.inbox", tx), nil
	})
	container.AddScoped("shoppingLists", func(c di.Container) (any, error) {
		return postgres.NewShoppingListRepository("depot.shopping_lists", c.Get("tx").(*sql.Tx)), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			"depot.stores_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewStoreRepository(svc.Config().Rpc.Address()),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			"depot.products_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewProductRepository(svc.Config().Rpc.Address()),
		), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return application.New(
			c.Get("shoppingLists").(domain.ShoppingListRepository),
			c.Get("stores").(domain.StoreCacheRepository),
			c.Get("products").(domain.ProductCacheRepository),
			c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent]),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get("eventPublisher").(am.EventPublisher)), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get("stores").(domain.StoreCacheRepository),
			c.Get("products").(domain.ProductCacheRepository),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return handlers.NewCommandHandlers(c.Get("app").(application.App)), nil
	})

	// setup Driver adapters
	if err := grpc.RegisterServerTx(container, svc.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, svc.Mux(), svc.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(svc.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(container)
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return nil
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("depot outbox processor encountered an error")
		}
	}()
}
