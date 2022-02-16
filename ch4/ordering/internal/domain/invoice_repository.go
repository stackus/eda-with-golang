package domain

import (
	"context"
)

type InvoiceRepository interface {
	Save(ctx context.Context, orderID OrderID, amount float64) (InvoiceID, error)
	Update(ctx context.Context, order *Order) error
	Delete(ctx context.Context, invoiceID InvoiceID) error
}
