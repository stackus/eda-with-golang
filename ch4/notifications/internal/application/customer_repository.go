package application

import (
	"context"

	"eda-in-golang/ch4/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
