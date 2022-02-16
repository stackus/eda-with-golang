package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/depot/depotpb"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type ShoppingListRepository struct {
	client depotpb.DepotServiceClient
}

var _ domain.ShoppingListRepository = (*ShoppingListRepository)(nil)

func NewShoppingListRepository(conn *grpc.ClientConn) ShoppingListRepository {
	return ShoppingListRepository{client: depotpb.NewDepotServiceClient(conn)}
}

func (r ShoppingListRepository) Save(ctx context.Context, order *domain.Order) error {
	// TODO implement me
	panic("implement me")
}

func (r ShoppingListRepository) Delete(ctx context.Context, orderID domain.OrderID) error {
	// TODO implement me
	panic("implement me")
}
