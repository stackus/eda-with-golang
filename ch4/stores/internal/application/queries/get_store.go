package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	repo ports.StoreRepository
}

func NewGetStoreHandler(repo ports.StoreRepository) GetStoreHandler {
	return GetStoreHandler{repo: repo}
}

func (h GetStoreHandler) Handle(ctx context.Context, query GetStore) (*domain.Store, error) {
	return h.repo.FindStore(ctx, query.ID)
}
