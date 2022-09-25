package cosec

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/cosec/internal"
	"eda-in-golang/cosec/internal/handlers"
	"eda-in-golang/cosec/internal/models"
	"eda-in-golang/customers/customerspb"
	"eda-in-golang/depot/depotpb"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/jetstream"
	pg "eda-in-golang/internal/postgres"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/registry/serdes"
	"eda-in-golang/internal/sec"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/tm"
	"eda-in-golang/ordering/orderingpb"
	"eda-in-golang/payments/paymentspb"
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
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := paymentspb.Registrations(reg); err != nil {
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
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.MessageStream),
			pg.NewOutboxStore("cosec.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddScoped("messagePublisher", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("cosec.outbox", tx)
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
		return pg.NewInboxStore("cosec.inbox", tx), nil
	})
	container.AddScoped("sagaRepo", func(c di.Container) (any, error) {
		reg := c.Get("registry").(registry.Registry)
		return sec.NewSagaRepository[*models.CreateOrderData](
			reg,
			pg.NewSagaStore(
				"cosec.sagas",
				c.Get("tx").(*sql.Tx),
				reg,
			),
		), nil
	})
	container.AddSingleton("saga", func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	container.AddScoped("orchestrator", func(c di.Container) (any, error) {
		return sec.NewOrchestrator[*models.CreateOrderData](
			c.Get("saga").(sec.Saga[*models.CreateOrderData]),
			c.Get("sagaRepo").(sec.SagaRepository[*models.CreateOrderData]),
			c.Get("commandPublisher").(am.CommandPublisher),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get("orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
		), nil
	})

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
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
			logger.Error().Err(err).Msg("cosec outbox processor encountered an error")
		}
	}()
}
