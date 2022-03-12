package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
