package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/depot/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch6/depot/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch6/depot/internal/domain"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingList) error
		CancelShoppingList(ctx context.Context, cmd commands.CancelShoppingList) error
		AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingList) error
		CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingList) error
	}
	Queries interface {
		GetShoppingList(ctx context.Context, query queries.GetShoppingList) (*domain.ShoppingList, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateShoppingListHandler
		commands.CancelShoppingListHandler
		commands.AssignShoppingListHandler
		commands.CompleteShoppingListHandler
	}
	appQueries struct {
		queries.GetShoppingListHandler
	}
)

var _ App = (*Application)(nil)

func New(shoppingLists domain.ShoppingListRepository, stores domain.StoreRepository, products domain.ProductRepository,
	domainPublisher ddd.EventPublisher,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateShoppingListHandler:   commands.NewCreateShoppingListHandler(shoppingLists, stores, products, domainPublisher),
			CancelShoppingListHandler:   commands.NewCancelShoppingListHandler(shoppingLists, domainPublisher),
			AssignShoppingListHandler:   commands.NewAssignShoppingListHandler(shoppingLists, domainPublisher),
			CompleteShoppingListHandler: commands.NewCompleteShoppingListHandler(shoppingLists, domainPublisher),
		},
		appQueries: appQueries{
			GetShoppingListHandler: queries.NewGetShoppingListHandler(shoppingLists),
		},
	}
}
