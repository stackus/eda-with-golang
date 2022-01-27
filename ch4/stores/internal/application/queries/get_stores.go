package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetStores struct {
}

type GetStoresHandler struct {
	repo ports.StoreRepository
}

func NewGetStoresHandler(repo ports.StoreRepository) GetStoresHandler {
	return GetStoresHandler{repo: repo}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.Store, error) {
	return h.repo.FindStores(ctx)
}
