package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type StoreRepository interface {
	Save(ctx context.Context, store *domain.Store) error
	Update(ctx context.Context, store *domain.Store) error
	Delete(ctx context.Context, storeID string) error
	Find(ctx context.Context, storeID string) (*domain.Store, error)
	FindAll(ctx context.Context) ([]*domain.Store, error)
}
