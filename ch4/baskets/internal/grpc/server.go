package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/baskets/basketspb"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application/commands"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application/queries"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
)

type server struct {
	app application.App
	basketspb.UnimplementedBasketServiceServer
}

var _ basketspb.BasketServiceServer = (*server)(nil)

func (s server) StartBasket(ctx context.Context, _ *basketspb.StartBasketRequest) (*basketspb.StartBasketResponse, error) {
	basketID := uuid.New().String()
	err := s.app.StartBasket(ctx, commands.StartBasket{
		ID: domain.BasketID(basketID),
	})

	return &basketspb.StartBasketResponse{Id: basketID}, err
}

func (s server) CancelBasket(ctx context.Context, request *basketspb.CancelBasketRequest) (*basketspb.CancelBasketResponse, error) {
	err := s.app.CancelBasket(ctx, commands.CancelBasket{
		ID: domain.BasketID(request.GetId()),
	})

	return &basketspb.CancelBasketResponse{}, err
}

func (s server) CheckoutBasket(ctx context.Context, request *basketspb.CheckoutBasketRequest) (*basketspb.CheckoutBasketResponse, error) {
	err := s.app.CheckoutBasket(ctx, commands.CheckoutBasket{
		ID:        domain.BasketID(request.GetId()),
		CardToken: request.GetCardToken(),
		SmsNumber: request.GetSmsNumber(),
	})

	return &basketspb.CheckoutBasketResponse{}, err
}

func (s server) AddItem(ctx context.Context, request *basketspb.AddItemRequest) (*basketspb.AddItemResponse, error) {
	err := s.app.AddItem(ctx, commands.AddItem{
		ID:        domain.BasketID(request.GetId()),
		ProductID: domain.ProductID(request.GetProductId()),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketspb.AddItemResponse{}, err
}

func (s server) RemoveItem(ctx context.Context, request *basketspb.RemoveItemRequest) (*basketspb.RemoveItemResponse, error) {
	err := s.app.RemoveItem(ctx, commands.RemoveItem{
		ID:        domain.BasketID(request.GetId()),
		ProductID: domain.ProductID(request.GetProductId()),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketspb.RemoveItemResponse{}, err
}

func (s server) GetBasket(ctx context.Context, request *basketspb.GetBasketRequest) (*basketspb.GetBasketResponse, error) {
	basket, err := s.app.GetBasket(ctx, queries.GetBasket{
		ID: domain.BasketID(request.GetId()),
	})
	if err != nil {
		return nil, err
	}

	return &basketspb.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}, nil
}

func (s server) basketFromDomain(basket *domain.Basket) *basketspb.Basket {
	protoBasket := &basketspb.Basket{
		Id: basket.ID.String(),
	}

	protoBasket.Items = make([]*basketspb.Item, 0, len(basket.Items))

	for _, item := range basket.Items {
		protoBasket.Items = append(protoBasket.Items, &basketspb.Item{
			StoreId:      item.StoreID.String(),
			StoreName:    item.StoreName,
			ProductId:    item.ProductID.String(),
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice.Float64(),
			Quantity:     int32(item.Quantity),
		})
	}

	return protoBasket
}
