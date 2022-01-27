package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/basket/basketpb"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/domain"
)

type server struct {
	app application.App
	basketpb.UnimplementedBasketServiceServer
}

var _ basketpb.BasketServiceServer = (*server)(nil)

func (s server) StartBasket(ctx context.Context, request *basketpb.StartBasketRequest) (*basketpb.StartBasketResponse, error) {
	basketID := uuid.New().String()
	err := s.app.StartBasket(ctx, commands.StartBasket{ID: basketID})

	return &basketpb.StartBasketResponse{Id: basketID}, err
}

func (s server) CancelBasket(ctx context.Context, request *basketpb.CancelBasketRequest) (*basketpb.CancelBasketResponse, error) {
	err := s.app.CancelBasket(ctx, commands.CancelBasket{ID: request.GetId()})

	return &basketpb.CancelBasketResponse{}, err
}

func (s server) CheckoutBasket(ctx context.Context, request *basketpb.CheckoutBasketRequest) (*basketpb.CheckoutBasketResponse, error) {
	err := s.app.CheckoutBasket(ctx, commands.CheckoutBasket{
		ID:        request.GetId(),
		CardToken: request.GetCardToken(),
		SmsNumber: request.GetSmsNumber(),
	})

	return &basketpb.CheckoutBasketResponse{}, err
}

func (s server) AddItem(ctx context.Context, request *basketpb.AddItemRequest) (*basketpb.AddItemResponse, error) {
	err := s.app.AddItem(ctx, commands.AddItem{
		ID:        request.GetId(),
		StoreID:   request.GetStoreId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketpb.AddItemResponse{}, err
}

func (s server) RemoveItem(ctx context.Context, request *basketpb.RemoveItemRequest) (*basketpb.RemoveItemResponse, error) {
	err := s.app.RemoveItem(ctx, commands.RemoveItem{
		ID:        request.GetId(),
		StoreID:   request.GetStoreId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketpb.RemoveItemResponse{}, err
}

func (s server) GetBasket(ctx context.Context, request *basketpb.GetBasketRequest) (*basketpb.GetBasketResponse, error) {
	basket, err := s.app.GetBasket(ctx, queries.GetBasket{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &basketpb.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}, nil
}

func (s server) basketFromDomain(basket *domain.Basket) *basketpb.Basket {
	protoBasket := &basketpb.Basket{
		Id: basket.ID,
	}

	protoBasket.Items = make([]*basketpb.Item, 0, len(basket.Items))

	for _, item := range basket.Items {
		protoBasket.Items = append(protoBasket.Items, &basketpb.Item{
			StoreId:      item.StoreID,
			ProductId:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     int32(item.Quantity),
		})
	}

	return protoBasket
}
