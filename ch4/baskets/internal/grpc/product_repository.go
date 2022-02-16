package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/stores/storespb"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type ProductRepository struct {
	client storespb.StoresServiceClient
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(conn *grpc.ClientConn) ProductRepository {
	return ProductRepository{client: storespb.NewStoresServiceClient(conn)}
}

func (r ProductRepository) Find(ctx context.Context, productID domain.ProductID) (*domain.Product, error) {
	resp, err := r.client.GetProduct(ctx, &storespb.GetProductRequest{
		Id: productID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "requesting product")
	}

	return r.productToDomain(resp.Product), nil
}

func (r ProductRepository) productToDomain(product *storespb.Product) *domain.Product {
	return &domain.Product{
		ID:      domain.ProductID(product.GetId()),
		StoreID: domain.StoreID(product.GetStoreId()),
		Name:    product.GetName(),
		Price:   domain.Price(product.GetPrice()),
	}
}
