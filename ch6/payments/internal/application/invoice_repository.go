package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/payments/internal/models"
)

type InvoiceRepository interface {
	Find(ctx context.Context, invoiceID string) (*models.Invoice, error)
	Save(ctx context.Context, invoice *models.Invoice) error
	Update(ctx context.Context, invoice *models.Invoice) error
}
