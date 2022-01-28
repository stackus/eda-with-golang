package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type ParticipatingStoreRepository interface {
	FindAll(ctx context.Context) ([]*domain.Store, error)
}
