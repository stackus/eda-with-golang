package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
		EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error
		DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error
		AddOffering(ctx context.Context, cmd commands.AddOffering) error
		RemoveOffering(ctx context.Context, cmd commands.RemoveOffering) error
	}
	Queries interface {
		GetStore(ctx context.Context, query queries.GetStore) (*domain.Store, error)
		GetStores(ctx context.Context, query queries.GetStores) ([]*domain.Store, error)
		GetParticipatingStores(ctx context.Context, query queries.GetParticipatingStores) ([]*domain.Store, error)
		GetStoreOfferings(ctx context.Context, query queries.GetStoreOfferings) ([]*domain.Offering, error)
		GetOffering(ctx context.Context, query queries.GetOffering) (*domain.Offering, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateStoreHandler
		commands.EnableParticipationHandler
		commands.DisableParticipationHandler
		commands.AddOfferingHandler
		commands.RemoveOfferingHandler
	}
	appQueries struct {
		queries.GetStoreHandler
		queries.GetStoresHandler
		queries.GetParticipatingStoresHandler
		queries.GetStoreOfferingsHandler
		queries.GetOfferingHandler
	}
)

var _ App = (*Application)(nil)

func New(storeRepo ports.StoreRepository, participatingStoreRepo ports.ParticipatingStoreRepository, offeringRepo ports.OfferingRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler:          commands.NewCreateStoreHandler(storeRepo),
			EnableParticipationHandler:  commands.NewEnableParticipationHandler(storeRepo),
			DisableParticipationHandler: commands.NewDisableParticipationHandler(storeRepo),
			AddOfferingHandler:          commands.NewAddOfferingHandler(storeRepo, offeringRepo),
			RemoveOfferingHandler:       commands.NewRemoveOfferingHandler(storeRepo, offeringRepo),
		},
		appQueries: appQueries{
			GetStoreHandler:               queries.NewGetStoreHandler(storeRepo),
			GetStoresHandler:              queries.NewGetStoresHandler(storeRepo),
			GetParticipatingStoresHandler: queries.NewGetParticipatingStoresHandler(participatingStoreRepo),
			GetStoreOfferingsHandler:      queries.NewGetStoreOfferingsHandler(offeringRepo),
			GetOfferingHandler:            queries.NewGetOfferingHandler(offeringRepo),
		},
	}
}
