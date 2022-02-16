package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	stores domain.StoreRepository
}

func NewGetStoreHandler(stores domain.StoreRepository) GetStoreHandler {
	return GetStoreHandler{stores: stores}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*domain.Store, error) {
	return h.stores.Find(ctx, query.ID)
}
