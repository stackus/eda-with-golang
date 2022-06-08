package customers

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"eda-in-golang/customers/customerspb"
	"eda-in-golang/customers/internal/application"
	"eda-in-golang/customers/internal/domain"
	"eda-in-golang/customers/internal/grpc"
	"eda-in-golang/customers/internal/handlers"
	"eda-in-golang/customers/internal/logging"
	"eda-in-golang/customers/internal/postgres"
	"eda-in-golang/customers/internal/rest"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/jetstream"
	"eda-in-golang/internal/monolith"
	"eda-in-golang/internal/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err = customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	// setup Driven adapters
	// reg := registry.New()
	// if err = customerspb.Registrations(reg); err != nil {
	// 	return err
	// }
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()), nil
	})
	// stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())

	container.AddSingleton("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("stream").(am.RawMessageStream)), nil
	})
	// eventStream := am.NewEventStream(reg, stream)

	container.AddSingleton("commandStream", func(c di.Container) (any, error) {
		return am.NewCommandStream(c.Get("registry").(registry.Registry), c.Get("stream").(am.RawMessageStream)), nil
	})
	// commandStream := am.NewCommandStream(reg, stream)

	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	// domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)

		return db.Begin()
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return postgres.NewCustomerRepository("customers.customers", c.Get("tx").(*sql.Tx)), nil
	})
	// customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.LogApplicationAccess(
			application.New(
				c.Get("customers").(domain.CustomerRepository),
				c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent]),
			),
			c.Get("logger").(zerolog.Logger),
		), nil
	})
	// app := logging.LogApplicationAccess(
	// 	application.New(customers, domainDispatcher),
	// 	mono.Logger(),
	// )

	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.AggregateEvent](
			// TODO this will need a scoped outboxEventStream
			handlers.NewDomainEventHandlers(container.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})
	// domainEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
	// 	handlers.NewDomainEventHandlers(eventStream),
	// 	"DomainEvents", mono.Logger(),
	// )

	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.LogCommandHandlerAccess[ddd.Command](
			handlers.NewCommandHandlers(c.Get("app").(application.App)),
			"Commands", c.Get("logger").(zerolog.Logger),
		), nil
	})
	// commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
	// 	handlers.NewCommandHandlers(app),
	// 	"Commands", mono.Logger(),
	// )

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(container.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent]))
	if err = handlers.RegisterCommandHandlersTx(container.Get("commandStream").(am.CommandStream), container); err != nil {
		return err
	}

	return nil
}
