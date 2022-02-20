package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
