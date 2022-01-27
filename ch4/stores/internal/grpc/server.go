package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/stores/storespb"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type server struct {
	app application.App
	storespb.UnimplementedStoresServiceServer
}

var _ storespb.StoresServiceServer = (*server)(nil)

func (s server) CreateStore(ctx context.Context, request *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
	storeID := uuid.New().String()

	err := s.app.CreateStore(ctx, commands.CreateStore{
		ID:       storeID,
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
	store, err := s.app.GetStore(ctx, queries.GetStore{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &storespb.GetStoreResponse{Store: s.storeFromDomain(store)}, nil
}

func (s server) GetStores(ctx context.Context, request *storespb.GetStoresRequest) (*storespb.GetStoresResponse, error) {
	stores, err := s.app.GetStores(ctx, queries.GetStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*storespb.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &storespb.GetStoresResponse{
		Stores: protoStores,
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

func (s server) DisableParticipation(ctx context.Context, request *storespb.DisableParticipationRequest) (*storespb.DisableParticipationResponse, error) {
	err := s.app.DisableParticipation(ctx, commands.DisableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.DisableParticipationResponse{}, nil
}

func (s server) GetParticipatingStores(ctx context.Context, request *storespb.GetParticipatingStoresRequest) (*storespb.GetParticipatingStoresResponse, error) {
	stores, err := s.app.GetParticipatingStores(ctx, queries.GetParticipatingStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*storespb.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &storespb.GetParticipatingStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s server) AddOffering(ctx context.Context, request *storespb.AddOfferingRequest) (*storespb.AddOfferingResponse, error) {
	id := uuid.New().String()
	err := s.app.AddOffering(ctx, commands.AddOffering{
		ID:          id,
		StoreID:     request.GetStoreId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		Price:       request.GetPrice(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.AddOfferingResponse{Id: id}, nil
}

func (s server) RemoveOffering(ctx context.Context, request *storespb.RemoveOfferingRequest) (*storespb.RemoveOfferingResponse, error) {
	err := s.app.RemoveOffering(ctx, commands.RemoveOffering{
		ID:      request.GetId(),
		StoreID: request.GetStoreId(),
	})

	return &storespb.RemoveOfferingResponse{}, err
}

func (s server) GetStoreOfferings(ctx context.Context, request *storespb.GetStoreOfferingsRequest) (*storespb.GetStoreOfferingsResponse, error) {
	offerings, err := s.app.GetStoreOfferings(ctx, queries.GetStoreOfferings{StoreID: request.GetStoreId()})
	if err != nil {
		return nil, err
	}

	protoOfferings := []*storespb.Offering{}
	for _, offering := range offerings {
		protoOfferings = append(protoOfferings, s.offeringFromDomain(offering))
	}

	return &storespb.GetStoreOfferingsResponse{
		Offerings: protoOfferings,
	}, nil
}

func (s server) GetOffering(ctx context.Context, request *storespb.GetOfferingRequest) (*storespb.GetOfferingResponse, error) {
	offering, err := s.app.GetOffering(ctx, queries.GetOffering{
		ID:      request.GetId(),
		StoreID: request.GetStoreId(),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.GetOfferingResponse{Offering: s.offeringFromDomain(offering)}, nil
}

func (s server) storeFromDomain(store *domain.Store) *storespb.Store {
	return &storespb.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location.String(),
		Participating: store.Participating,
	}
}

func (s server) offeringFromDomain(offering *domain.Offering) *storespb.Offering {
	return &storespb.Offering{
		Id:          offering.ID,
		StoreId:     offering.StoreID,
		Name:        offering.Name,
		Description: offering.Description,
		Price:       offering.Price,
	}
}
