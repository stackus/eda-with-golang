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

func (r InvoiceRepository) Save(ctx context.Context, orderID, paymentID string, amount float64) (string, error) {
	resp, err := r.client.CreateInvoice(ctx, &paymentspb.CreateInvoiceRequest{
		OrderId:   orderID,
		PaymentId: paymentID,
		Amount:    amount,
	})
	if err != nil {
		return "", err
	}

	return resp.GetId(), nil
}

func (r InvoiceRepository) Delete(ctx context.Context, invoiceID string) error {
	_, err := r.client.CancelInvoice(ctx, &paymentspb.CancelInvoiceRequest{Id: invoiceID})
	return err
}
