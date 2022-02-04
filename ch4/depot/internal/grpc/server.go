package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/depot/depotpb"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/application"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/application/commands"
)

type server struct {
	app application.App
	depotpb.UnimplementedDepotServiceServer
}

var _ depotpb.DepotServiceServer = (*server)(nil)

func (s server) SubmitOrder(ctx context.Context, request *depotpb.SubmitOrderRequest) (*depotpb.SubmitOrderResponse, error) {
	id := uuid.New().String()

	items := make([]commands.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.BuildShoppingList(ctx, commands.BuildShoppingList{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	return &depotpb.SubmitOrderResponse{Id: id}, err
}

func (s server) CancelOrder(ctx context.Context, request *depotpb.CancelOrderRequest) (*depotpb.CancelOrderResponse, error) {
	err := s.app.CancelOrder(ctx, commands.CancelOrder{OrderID: request.GetId()})

	return &depotpb.CancelOrderResponse{}, err
}

func (s server) itemToDomain(item *depotpb.OrderItem) commands.OrderItem {
	return commands.OrderItem{
		StoreID:   item.GetStoreId(),
		ProductID: item.GetProductId(),
		Quantity:  int(item.GetQuantity()),
	}
}
