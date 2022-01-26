package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type App interface {
	CreateStore(ctx context.Context, cmd commands.CreateStore) (string, error)
	EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error

	GetStore(ctx context.Context, query queries.GetStore) (*domain.Store, error)
}

type Application struct {
	c Commands
	q Queries
}

type Commands struct {
	createStore         commands.CreateStoreHandler
	enableParticipation commands.EnableParticipationHandler
}

type Queries struct {
	getStore queries.GetStoreHandler
}

var _ App = (*Application)(nil)

func New(storeRepo ports.StoreRepository) *Application {
	return &Application{
		c: Commands{
			createStore:         commands.NewCreateStoreHandler(storeRepo),
			enableParticipation: commands.NewEnableParticipationHandler(storeRepo),
		},
		q: Queries{
			getStore: queries.NewGetStoreHandler(storeRepo),
		},
	}
}

func (a Application) CreateStore(ctx context.Context, cmd commands.CreateStore) (string, error) {
	return a.c.createStore.Handle(ctx, cmd)
}

func (a Application) EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error {
	return a.c.enableParticipation.Handle(ctx, cmd)
}

func (a Application) GetStore(ctx context.Context, query queries.GetStore) (*domain.Store, error) {
	return a.q.getStore.Handle(ctx, query)
}
