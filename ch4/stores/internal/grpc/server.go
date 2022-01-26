package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/proto/storespb"
)

type server struct {
	app application.App
	storespb.UnimplementedStoresServiceServer
}

var _ storespb.StoresServiceServer = (*server)(nil)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	storespb.RegisterStoresServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateStore(ctx context.Context, request *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
	storeID, err := s.app.CreateStore(ctx, commands.CreateStore{
		Name:     request.GetName(),
		Location: request.GetLocation(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.CreateStoreResponse{
		Id: storeID,
	}, nil
}

func (s server) GetStore(ctx context.Context, request *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
	store, err := s.app.GetStore(ctx, queries.GetStore{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.GetStoreResponse{
		Store: s.storeFromDomain(store),
	}, nil
}

func (s server) EnableParticipation(ctx context.Context, request *storespb.EnableParticipationRequest) (*storespb.EnableParticipationResponse, error) {
	err := s.app.EnableParticipation(ctx, commands.EnableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.EnableParticipationResponse{}, nil
}

func (s server) storeFromDomain(store *domain.Store) *storespb.Store {
	return &storespb.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location.String(),
		Participating: store.Participating,
	}
}
