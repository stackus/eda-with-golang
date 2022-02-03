package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(orderRepo domain.OrderRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(orderRepo),
		},
		appQueries: appQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orderRepo),
		},
	}
}
