package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch5/ordering/internal/domain"
	"github.com/stackus/eda-with-golang/ch5/payments/paymentspb"
)

type PaymentRepository struct {
	client paymentspb.PaymentsServiceClient
}

var _ domain.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(conn *grpc.ClientConn) PaymentRepository {
	return PaymentRepository{
		client: paymentspb.NewPaymentsServiceClient(conn),
	}
}

func (r PaymentRepository) Confirm(ctx context.Context, paymentID string) error {
	_, err := r.client.ConfirmPayment(ctx, &paymentspb.ConfirmPaymentRequest{Id: paymentID})
	return err
}
