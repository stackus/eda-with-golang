package search

import (
	"context"
	"database/sql"

	"eda-in-golang/customers/customerspb"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/amotel"
	"eda-in-golang/internal/amprom"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/jetstream"
	pg "eda-in-golang/internal/postgres"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/tm"
	"eda-in-golang/ordering/orderingpb"
	"eda-in-golang/search/internal/application"
	"eda-in-golang/search/internal/grpc"
	"eda-in-golang/search/internal/handlers"
	"eda-in-golang/search/internal/postgres"
	"eda-in-golang/search/internal/rest"
	"eda-in-golang/stores/storespb"
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
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	container.AddScoped("tx", func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddSingleton("messageSubscriber", func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter("search"),
		), nil
	})
	container.AddScoped("inboxStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		return pg.NewInboxStore("search.inbox", tx), nil
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return postgres.NewCustomerCacheRepository(
			"search.customers_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewCustomerRepository(svc.Config().Rpc.Service("CUSTOMERS")),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			"search.stores_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewStoreRepository(svc.Config().Rpc.Service("STORES")),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			"search.products_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewProductRepository(svc.Config().Rpc.Service("STORES")),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return postgres.NewOrderRepository("search.orders", c.Get("tx").(*sql.Tx)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return application.New(
			c.Get("orders").(application.OrderRepository),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get("registry").(registry.Registry),
			c.Get("orders").(application.OrderRepository),
			c.Get("customers").(application.CustomerCacheRepository),
			c.Get("stores").(application.StoreCacheRepository),
			c.Get("products").(application.ProductCacheRepository),
			tm.InboxHandler(c.Get("inboxStore").(tm.InboxStore)),
		), nil
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

	return nil
}
