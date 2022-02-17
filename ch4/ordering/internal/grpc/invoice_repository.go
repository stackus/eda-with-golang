package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
	"github.com/stackus/eda-with-golang/ch4/payments/paymentspb"
)

type InvoiceRepository struct {
	client paymentspb.PaymentsServiceClient
}

var _ domain.InvoiceRepository = (*InvoiceRepository)(nil)

func NewInvoiceRepository(conn *grpc.ClientConn) InvoiceRepository {
	return InvoiceRepository{client: paymentspb.NewPaymentsServiceClient(conn)}
}

func (r InvoiceRepository) Save(ctx context.Context, orderID domain.OrderID, amount float64) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r InvoiceRepository) Update(ctx context.Context, invoiceID string, amount float64) error {
	// TODO implement me
	panic("implement me")
}

func (r InvoiceRepository) Delete(ctx context.Context, invoiceID string) error {
	// TODO implement me
	panic("implement me")
}
