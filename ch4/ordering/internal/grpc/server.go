package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
	"github.com/stackus/eda-with-golang/ch4/ordering/orderingpb"
)

type server struct {
	app application.App
	orderingpb.UnimplementedOrderingServiceServer
}

var _ orderingpb.OrderingServiceServer = (*server)(nil)

func (s server) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (*orderingpb.CreateOrderResponse, error) {
	id := uuid.New().String()

	items := make([]*domain.Item, 0, len(request.Items))
	for _, item := range request.Items {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateOrder(ctx, commands.CreateOrder{
		ID:        id,
		Items:     items,
		CardToken: request.CardToken,
		SmsNumber: request.SmsNumber,
	})

	return &orderingpb.CreateOrderResponse{Id: id}, err
}

func (s server) GetOrder(ctx context.Context, request *orderingpb.GetOrderRequest) (*orderingpb.GetOrderResponse, error) {
	order, err := s.app.GetOrder(ctx, queries.GetOrder{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &orderingpb.GetOrderResponse{
		Order: s.orderFromDomain(order),
	}, nil
}

func (s server) orderFromDomain(order *domain.Order) *orderingpb.Order {
	items := make([]*orderingpb.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, s.itemFromDomain(item))
	}

	return &orderingpb.Order{
		Id:        order.ID,
		Items:     items,
		CardToken: order.CardToken,
		SmsNumber: order.SmsNumber,
		Status:    order.Status.String(),
	}
}

func (s server) itemToDomain(item *orderingpb.Item) *domain.Item {
	return &domain.Item{
		ProductID: item.GetProductId(),
		StoreID:   item.GetStoreId(),
		Price:     item.GetPrice(),
		Quantity:  int(item.GetQuantity()),
	}
}

func (s server) itemFromDomain(item *domain.Item) *orderingpb.Item {
	return &orderingpb.Item{
		StoreId:   item.StoreID,
		ProductId: item.ProductID,
		Price:     item.Price,
		Quantity:  int32(item.Quantity),
	}
}
