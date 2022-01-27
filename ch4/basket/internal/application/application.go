package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
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

func New(basketRepo ports.BasketRepository, productRepo ports.ProductRepository) *Application {
	return &Application{
		appCommands: appCommands{
			StartBasketHandler:    commands.NewStartBasketHandler(basketRepo),
			CancelBasketHandler:   commands.NewCancelBasketHandler(basketRepo),
			CheckoutBasketHandler: commands.NewCheckoutBasketHandler(basketRepo),
			AddItemHandler:        commands.NewAddItemHandler(basketRepo, productRepo),
			RemoveItemHandler:     commands.NewRemoveItemHandler(basketRepo, productRepo),
		},
		appQueries: appQueries{
			GetBasketHandler: queries.NewGetBasketHandler(basketRepo),
		},
	}
}
