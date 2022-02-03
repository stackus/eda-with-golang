package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		BuildShoppingList(ctx context.Context, cmd commands.BuildShoppingList) error
		CancelOrder(ctx context.Context, cmd commands.CancelOrder) error
	}
	Queries interface {
		GetShoppingList(ctx context.Context, list queries.GetShoppingList) (*domain.ShoppingList, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.BuildShoppingListHandler
		commands.CancelOrderHandler
	}
	appQueries struct {
		queries.GetShoppingListHandler
	}
)

var _ App = (*Application)(nil)

func New(shoppingListRepo domain.ShoppingListRepository, storeRepo domain.StoreRepository, productRepo domain.ProductRepository) *Application {
	return &Application{
		appCommands: appCommands{
			BuildShoppingListHandler: commands.NewSubmitOrderHandler(shoppingListRepo, storeRepo, productRepo),
			CancelOrderHandler:       commands.NewCancelOrderHandler(shoppingListRepo),
		},
		appQueries: appQueries{
			GetShoppingListHandler: queries.NewGetShoppingListHandler(shoppingListRepo),
		},
	}
}
