package application

import (
	"context"

	"eda-in-golang/ch8/search/internal/models"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*models.Store, error)
}
