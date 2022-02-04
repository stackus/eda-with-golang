package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	repo domain.StoreRepository
}

func NewGetStoreHandler(repo domain.StoreRepository) GetStoreHandler {
	return GetStoreHandler{repo: repo}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*domain.Store, error) {
	return h.repo.Find(ctx, query.ID)
}
