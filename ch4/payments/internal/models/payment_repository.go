package models

import (
	"context"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *Payment) error
	Find(ctx context.Context, paymentID string) (*Payment, error)
}
