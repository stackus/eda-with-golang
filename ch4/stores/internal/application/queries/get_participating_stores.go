package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetParticipatingStores struct {
}

type GetParticipatingStoresHandler struct {
	repo ports.ParticipatingStoreRepository
}

func NewGetParticipatingStoresHandler(repo ports.ParticipatingStoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{repo: repo}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.Store, error) {
	return h.repo.FindAll(ctx)
}
