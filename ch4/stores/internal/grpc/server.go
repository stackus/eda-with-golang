package grpc

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	storespb "github.com/stackus/eda-with-golang/ch4/stores/internal/pb"
)

type server struct {
	app application.App
	storespb.UnimplementedStoresServiceServer
}

var _ storespb.StoresServiceServer = (*server)(nil)

func Register(app application.App, mono monolith.Monolith) error {
	storespb.RegisterStoresServiceServer(mono.RPC(), server{app: app})
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
	// TODO implement me
	panic("implement me")
}

func (s server) EnableParticipation(ctx context.Context, request *storespb.EnableParticipationRequest) (*storespb.EnableParticipationResponse, error) {
	// TODO implement me
	panic("implement me")
}
