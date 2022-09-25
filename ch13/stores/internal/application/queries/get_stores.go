package queries

import (
	"context"

	"eda-in-golang/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall domain.MallRepository
}

func NewGetStoresHandler(mall domain.MallRepository) GetStoresHandler {
	return GetStoresHandler{mall: mall}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.MallStore, error) {
	ctx, span := tracer.Start(ctx, "GetStores")
	defer span.End()

	return h.mall.All(ctx)
}
