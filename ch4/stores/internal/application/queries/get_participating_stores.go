package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetParticipatingStores struct {
}

type GetParticipatingStoresHandler struct {
	repo ports.StoreRepository
}

func NewGetParticipatingStoresHandler(repo ports.StoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{repo: repo}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, query GetParticipatingStores) ([]*domain.Store, error) {
	return h.repo.FindParticipatingStores(ctx)
}
