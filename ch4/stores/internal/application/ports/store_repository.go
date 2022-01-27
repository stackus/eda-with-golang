package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type StoreRepository interface {
	SaveStore(ctx context.Context, store *domain.Store) error
	UpdateStore(ctx context.Context, store *domain.Store) error
	DeleteStore(ctx context.Context, storeID string) error
	FindStore(ctx context.Context, storeID string) (*domain.Store, error)
	FindStores(ctx context.Context) ([]*domain.Store, error)
	FindParticipatingStores(ctx context.Context) ([]*domain.Store, error)
}
