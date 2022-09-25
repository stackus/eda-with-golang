package ordering

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/baskets/basketspb"
	"eda-in-golang/depot/depotpb"
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
	"eda-in-golang/ordering/internal/application"
	"eda-in-golang/ordering/internal/domain"
	"eda-in-golang/ordering/internal/grpc"
	"eda-in-golang/ordering/internal/handlers"
	"eda-in-golang/ordering/internal/rest"
	"eda-in-golang/ordering/orderingpb"
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
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := basketspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := orderingpb.Registrations(reg); err != nil {
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
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.MessageStream),
			pg.NewOutboxStore("ordering.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddScoped("messagePublisher", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("ordering.outbox", tx)
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
	container.AddScoped("replyPublisher", func(c di.Container) (any, error) {
		return am.NewReplyPublisher(
			c.Get("registry").(registry.Registry),
			c.Get("messagePublisher").(am.MessagePublisher),
		), nil
	})
	container.AddScoped("inboxStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		return pg.NewInboxStore("ordering.inbox", tx), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			pg.NewEventStore("ordering.events", tx, reg),
			pg.NewSnapshotStore("ordering.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Order](
			domain.OrderAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return application.New(
			c.Get("orders").(domain.OrderRepository),
			c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.Event]),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get("eventPublisher").(am.EventPublisher)), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get("app").(application.App),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return handlers.NewCommandHandlers(c.Get("app").(application.App)), nil
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

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err = serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}
	// order events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderRejected{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderApproved{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err = serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("ordering outbox processor encountered an error")
		}
	}()
}
