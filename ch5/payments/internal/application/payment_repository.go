package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/payments/internal/models"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *models.Payment) error
	Find(ctx context.Context, paymentID string) (*models.Payment, error)
}
