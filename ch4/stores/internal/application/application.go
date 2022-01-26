package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
)

type App interface {
	CreateStore(ctx context.Context, cmd commands.CreateStore) (string, error)
}

type Application struct {
	c Commands
}

type Commands struct {
	createStore commands.CreateStoreHandler
}

var _ App = (*Application)(nil)

func New(storeRepo ports.StoreRepository) *Application {
	return &Application{
		c: Commands{
			createStore: commands.NewCreateStoreHandler(storeRepo),
		},
	}
}

func (a Application) CreateStore(ctx context.Context, cmd commands.CreateStore) (string, error) {
	return a.c.createStore.Handle(ctx, cmd)
}
