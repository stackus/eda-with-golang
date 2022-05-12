package application

import (
	"context"

	"eda-in-golang/ch9/search/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
