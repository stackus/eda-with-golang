package domain

import (
	"context"
)

type InvoiceRepository interface {
	Save(ctx context.Context, orderID OrderID, amount float64) (string, error)
	Update(ctx context.Context, invoiceID string, amount float64) error
	Delete(ctx context.Context, invoiceID string) error
}
