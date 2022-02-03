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

func New(basketRepo domain.BasketRepository, productRepo domain.ProductRepository, orderRepo domain.OrderRepository) *Application {
	return &Application{
		appCommands: appCommands{
			StartBasketHandler:    commands.NewStartBasketHandler(basketRepo),
			CancelBasketHandler:   commands.NewCancelBasketHandler(basketRepo),
			CheckoutBasketHandler: commands.NewCheckoutBasketHandler(basketRepo, orderRepo),
			AddItemHandler:        commands.NewAddItemHandler(basketRepo, productRepo),
			RemoveItemHandler:     commands.NewRemoveItemHandler(basketRepo, productRepo),
		},
		appQueries: appQueries{
			GetBasketHandler: queries.NewGetBasketHandler(basketRepo),
		},
	}
}
