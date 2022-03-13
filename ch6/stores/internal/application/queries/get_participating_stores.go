package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type GetParticipatingStores struct{}

type GetParticipatingStoresHandler struct {
	mall domain.MallRepository
}

func NewGetParticipatingStoresHandler(mall domain.MallRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{mall: mall}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores,
) ([]*domain.Store, error) {
	return h.mall.AllParticipating(ctx)
}
