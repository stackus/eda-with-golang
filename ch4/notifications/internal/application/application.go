package application

import (
	"context"
)

type (
	OrderCreated struct {
		SMSNumber string
		OrderID   string
	}

	App interface {
		NotifyOrderCreated(ctx context.Context, notify OrderCreated) error
	}

	Application struct {
	}
)

var _ App = (*Application)(nil)

func New() *Application {
	return &Application{}
}

func (a Application) NotifyOrderCreated(ctx context.Context, notify OrderCreated) error {
	// not implemented

	return nil
}
