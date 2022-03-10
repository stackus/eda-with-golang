package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch5/ordering/orderingpb"
	"github.com/stackus/eda-with-golang/ch5/payments/internal/application"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Complete(ctx context.Context, invoiceID, orderID string) error {
	_, err := r.client.CompleteOrder(ctx, &orderingpb.CompleteOrderRequest{
		Id:        orderID,
		InvoiceId: invoiceID,
	})
	return err
}
