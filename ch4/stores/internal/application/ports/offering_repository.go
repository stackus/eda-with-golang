package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type OfferingRepository interface {
	FindOffering(ctx context.Context, id, storeID string) (*domain.Offering, error)
	AddOffering(ctx context.Context, offering *domain.Offering) error
	RemoveOffering(ctx context.Context, id, storeID string) error
	GetStoreOfferings(ctx context.Context, storeID string) ([]*domain.Offering, error)
}
