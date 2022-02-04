package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
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

func New(storeRepo domain.StoreRepository, participatingStoreRepo domain.ParticipatingStoreRepository, productRepo domain.ProductRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler:          commands.NewCreateStoreHandler(storeRepo),
			EnableParticipationHandler:  commands.NewEnableParticipationHandler(storeRepo),
			DisableParticipationHandler: commands.NewDisableParticipationHandler(storeRepo),
			AddProductHandler:           commands.NewAddProductHandler(storeRepo, productRepo),
			RemoveProductHandler:        commands.NewRemoveProductHandler(storeRepo, productRepo),
		},
		appQueries: appQueries{
			GetStoreHandler:               queries.NewGetStoreHandler(storeRepo),
			GetStoresHandler:              queries.NewGetStoresHandler(storeRepo),
			GetParticipatingStoresHandler: queries.NewGetParticipatingStoresHandler(participatingStoreRepo),
			GetCatalogHandler:             queries.NewGetCatalogHandler(productRepo),
			GetProductHandler:             queries.NewGetProductHandler(productRepo),
		},
	}
}
