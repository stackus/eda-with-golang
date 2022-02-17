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
		CancelOrder(ctx context.Context, cmd commands.CancelOrder) error
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
		commands.CancelOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository, invoices domain.InvoiceRepository, shopping domain.ShoppingRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(orders, customers, payments, shopping),
			CancelOrderHandler: commands.NewCancelOrderHandler(orders, shopping),
		},
		appQueries: appQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orders),
		},
	}
}
