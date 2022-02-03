package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/baskets/internal/domain"
	"github.com/stackus/eda-with-golang/ch4/ordering/orderingpb"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{client: orderingpb.NewOrderingServiceClient(conn)}
}

func (r OrderRepository) Save(ctx context.Context, basket *domain.Basket) (string, error) {
	items := make([]*orderingpb.Item, 0, len(basket.Items))
	for _, item := range basket.Items {
		items = append(items, &orderingpb.Item{
			StoreId:   item.StoreID,
			ProductId: item.ProductID,
			Price:     item.ProductPrice,
			Quantity:  int32(item.Quantity),
		})
	}

	resp, err := r.client.CreateOrder(ctx, &orderingpb.CreateOrderRequest{
		Items:     items,
		CardToken: basket.CardToken,
		SmsNumber: basket.SmsNumber,
	})
	if err != nil {
		return "", errors.Wrap(err, "saving order")
	}

	return resp.GetId(), nil
}
