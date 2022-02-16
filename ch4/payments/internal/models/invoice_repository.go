package models

import (
	"context"
)

type InvoiceRepository interface {
	Find(ctx context.Context, invoiceID string) (*Invoice, error)
	Save(ctx context.Context, invoice *Invoice) error
	Update(ctx context.Context, invoice *Invoice) error
}
