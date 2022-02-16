package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/customers/internal/domain"
)

type (
	RegisterCustomer struct {
		ID domain.CustomerID
	}

	App interface {
		RegisterCustomer(ctx context.Context, register RegisterCustomer) error
	}

	Application struct{}
)

func New() *Application {
	return &Application{}
}
