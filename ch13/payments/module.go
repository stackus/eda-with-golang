package payments

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/amotel"
	"eda-in-golang/internal/amprom"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/jetstream"
	pg "eda-in-golang/internal/postgres"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/tm"
	"eda-in-golang/ordering/orderingpb"
	"eda-in-golang/payments/internal/application"
	"eda-in-golang/payments/internal/grpc"
	"eda-in-golang/payments/internal/handlers"
	"eda-in-golang/payments/internal/postgres"
	"eda-in-golang/payments/internal/rest"
	"eda-in-golang/payments/paymentspb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := orderingpb.Registrations(reg); err != nil {
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
	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.MessageStream),
			pg.NewOutboxStore("payments.outbox", svc.DB()),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddSingleton("sentCounter", func(c di.Container) (any, error) {
		return amprom.SentMessagesCounter("payments"), nil
	})
	container.AddScoped("messagePublisher", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("payments.outbox", tx)
		return am.NewMessagePublisher(
			c.Get("stream").(am.MessageStream),
			amotel.OtelMessageContextInjector(),
			c.Get("sentCounter").(am.MessagePublisherMiddleware),
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	container.AddSingleton("messageSubscriber", func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			c.Get("stream").(am.MessageStream),
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter("payments"),
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
		return pg.NewInboxStore("payments.inbox", tx), nil
	})
	container.AddScoped("invoices", func(c di.Container) (any, error) {
		return postgres.NewInvoiceRepository("payments.invoices", c.Get("tx").(*sql.Tx)), nil
	})
	container.AddScoped("payments", func(c di.Container) (any, error) {
		return postgres.NewPaymentRepository("payments.payments", c.Get("tx").(*sql.Tx)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return application.New(
			c.Get("invoices").(application.InvoiceRepository),
			c.Get("payments").(application.PaymentRepository),
			c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.Event]),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get("eventPublisher").(am.EventPublisher)), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewIntegrationHandlers(
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
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(container)
	if err = handlers.RegisterCommandHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("payments outbox processor encountered an error")
		}
	}()
}
