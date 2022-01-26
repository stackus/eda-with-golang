package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type StoreRepository interface {
	FindStore(ctx context.Context, storeID string) (*domain.Store, error)
	SaveStore(ctx context.Context, store *domain.Store) error
	UpdateStore(ctx context.Context, store *domain.Store) error
}
