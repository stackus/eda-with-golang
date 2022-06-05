package depot

import (
	"context"

	"eda-in-golang/depot/depotpb"
	"eda-in-golang/depot/internal/application"
	"eda-in-golang/depot/internal/grpc"
	"eda-in-golang/depot/internal/handlers"
	"eda-in-golang/depot/internal/logging"
	"eda-in-golang/depot/internal/postgres"
	"eda-in-golang/depot/internal/rest"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/jetstream"
	"eda-in-golang/internal/monolith"
	"eda-in-golang/internal/registry"
	"eda-in-golang/stores/storespb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// container := di.New()

	// setup Driven adapters
	// container.AddSingleton("registry", func(c di.Container) (any, error) {
	// 	reg := registry.New()
	// 	if err = storespb.Registrations(reg); err != nil {
	// 		return nil, err
	// 	}
	// 	if err = depotpb.Registrations(reg); err != nil {
	// 		return nil, err
	// 	}
	// 	return reg, nil
	// })
	reg := registry.New()
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	if err = depotpb.Registrations(reg); err != nil {
		return err
	}

	// container.AddSingleton("logger", func(c di.Container) (any, error) {
	// 	return mono.Logger(), nil
	// })
	// container.AddSingleton("stream", func(c di.Container) (any, error) {
	// 	return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), c.Get("logger").(zerolog.Logger)), nil
	// })
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())

	// outboxStream := outbox.NewStream(messageStore, stream)

	// container.AddSingleton("eventStream", func(c di.Container) (any, error) {
	// 	return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("stream").(am.RawMessageStream)), nil
	// })
	eventStream := am.NewEventStream(reg, stream) // (reg, outboxStream)

	// container.AddSingleton("commandStream", func(c di.Container) (any, error) {
	// 	return am.NewCommandStream(c.Get("registry").(registry.Registry), c.Get("stream").(am.RawMessageStream)), nil
	// })
	commandStream := am.NewCommandStream(reg, stream) // (reg, outboxStream)

	// container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
	// 	return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	// })
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

	// container.AddSingleton("db", func(c di.Container) (any, error) {
	// 	return mono.DB(), nil
	// })

	shoppingLists := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := postgres.NewStoreCacheRepository("depot.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
	products := postgres.NewProductCacheRepository("depot.products_cache", mono.DB(), grpc.NewProductRepository(conn))

	// setup application
	app := logging.LogApplicationAccess(
		// middleware that creates the repos+application using the transaction
		application.New(shoppingLists, stores, products, domainDispatcher),
		mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		// middleware that creates the repos+stream using the transaction
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		// middleware that creates the repos+handler using the transaction
		handlers.NewIntegrationEventHandlers(stores, products),
		"IntegrationEvents", mono.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		// instead of app this will accept the container?
		// or still takes the app but the app uses a middleware that starts a "scope"
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.Register(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

	return nil
}
