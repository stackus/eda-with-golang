package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/stores/internal/domain"
)

type GetStores struct {
}

type GetStoresHandler struct {
	stores domain.StoreRepository
}

func NewGetStoresHandler(stores domain.StoreRepository) GetStoresHandler {
	return GetStoresHandler{stores: stores}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.Store, error) {
	return h.stores.FindAll(ctx)
}
