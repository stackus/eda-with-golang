package application

import (
	"context"

	"eda-in-golang/ch9/search/internal/models"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*models.Store, error)
}
