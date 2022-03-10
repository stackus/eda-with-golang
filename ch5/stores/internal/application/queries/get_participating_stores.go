package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/stores/internal/domain"
)

type GetParticipatingStores struct {
}

type GetParticipatingStoresHandler struct {
	participatingStores domain.ParticipatingStoreRepository
}

func NewGetParticipatingStoresHandler(participatingStores domain.ParticipatingStoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{participatingStores: participatingStores}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores,
) ([]*domain.Store, error) {
	return h.participatingStores.FindAll(ctx)
}
