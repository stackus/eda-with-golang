package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/stores/storespb"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type ProductRepository struct {
	client storespb.StoresServiceClient
}

var _ ports.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(conn *grpc.ClientConn) ProductRepository {
	return ProductRepository{client: storespb.NewStoresServiceClient(conn)}
}

func (r ProductRepository) FindProduct(ctx context.Context, storeID, productID string) (*domain.Product, error) {
	resp, err := r.client.GetOffering(ctx, &storespb.GetOfferingRequest{
		Id:      productID,
		StoreId: storeID,
	})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(resp.Offering), nil
}

func (r ProductRepository) productToDomain(offering *storespb.Offering) *domain.Product {
	return &domain.Product{
		ID:      offering.GetId(),
		StoreID: offering.GetStoreId(),
		Name:    offering.GetName(),
		Price:   offering.GetPrice(),
	}
}
