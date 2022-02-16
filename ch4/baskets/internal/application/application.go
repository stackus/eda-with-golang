package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		StartBasket(ctx context.Context, cmd commands.StartBasket) error
		CancelBasket(ctx context.Context, cmd commands.CancelBasket) error
		CheckoutBasket(ctx context.Context, cmd commands.CheckoutBasket) error
		AddItem(ctx context.Context, cmd commands.AddItem) error
		RemoveItem(ctx context.Context, cmd commands.RemoveItem) error
	}
	Queries interface {
		GetBasket(ctx context.Context, query queries.GetBasket) (*domain.Basket, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.StartBasketHandler
		commands.CancelBasketHandler
		commands.CheckoutBasketHandler
		commands.AddItemHandler
		commands.RemoveItemHandler
	}
	appQueries struct {
		queries.GetBasketHandler
	}
)

var _ App = (*Application)(nil)

func New(baskets domain.BasketRepository, stores domain.StoreRepository, products domain.ProductRepository, orders domain.OrderRepository) *Application {
	return &Application{
		appCommands: appCommands{
			StartBasketHandler:    commands.NewStartBasketHandler(baskets),
			CancelBasketHandler:   commands.NewCancelBasketHandler(baskets),
			CheckoutBasketHandler: commands.NewCheckoutBasketHandler(baskets, orders),
			AddItemHandler:        commands.NewAddItemHandler(baskets, stores, products),
			RemoveItemHandler:     commands.NewRemoveItemHandler(baskets, products),
		},
		appQueries: appQueries{
			GetBasketHandler: queries.NewGetBasketHandler(baskets),
		},
	}
}
