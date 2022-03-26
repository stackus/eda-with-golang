package queries

import (
	"context"

	"eda-in-golang/ch7/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall domain.MallRepository
}

func NewGetStoresHandler(mall domain.MallRepository) GetStoresHandler {
	return GetStoresHandler{mall: mall}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.Store, error) {
	return h.mall.All(ctx)
}
