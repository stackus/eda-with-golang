package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
		EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error
		DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error
		RebrandStore(ctx context.Context, cmd commands.RebrandStore) error
		AddProduct(ctx context.Context, cmd commands.AddProduct) error
		RemoveProduct(ctx context.Context, cmd commands.RemoveProduct) error
	}
	Queries interface {
		GetStore(ctx context.Context, query queries.GetStore) (*domain.Store, error)
		GetStores(ctx context.Context, query queries.GetStores) ([]*domain.Store, error)
		GetParticipatingStores(ctx context.Context, query queries.GetParticipatingStores) ([]*domain.Store, error)
		GetCatalog(ctx context.Context, query queries.GetCatalog) ([]*domain.Product, error)
		GetProduct(ctx context.Context, query queries.GetProduct) (*domain.Product, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateStoreHandler
		commands.EnableParticipationHandler
		commands.DisableParticipationHandler
		commands.RebrandStoreHandler
		commands.AddProductHandler
		commands.RemoveProductHandler
	}
	appQueries struct {
		queries.GetStoreHandler
		queries.GetStoresHandler
		queries.GetParticipatingStoresHandler
		queries.GetCatalogHandler
		queries.GetProductHandler
	}
)

var _ App = (*Application)(nil)

func New(stores domain.StoreRepository, products domain.ProductRepository,
	catalog domain.CatalogRepository, mall domain.MallRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler:          commands.NewCreateStoreHandler(stores),
			EnableParticipationHandler:  commands.NewEnableParticipationHandler(stores),
			DisableParticipationHandler: commands.NewDisableParticipationHandler(stores),
			RebrandStoreHandler:         commands.NewRebrandStoreHandler(stores),
			AddProductHandler:           commands.NewAddProductHandler(stores, products),
			RemoveProductHandler:        commands.NewRemoveProductHandler(products),
		},
		appQueries: appQueries{
			GetStoreHandler:               queries.NewGetStoreHandler(mall),
			GetStoresHandler:              queries.NewGetStoresHandler(mall),
			GetParticipatingStoresHandler: queries.NewGetParticipatingStoresHandler(mall),
			GetCatalogHandler:             queries.NewGetCatalogHandler(catalog),
			GetProductHandler:             queries.NewGetProductHandler(catalog),
		},
	}
}
