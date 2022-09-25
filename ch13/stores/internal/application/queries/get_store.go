package queries

import (
	"context"

	"eda-in-golang/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	mall domain.MallRepository
}

func NewGetStoreHandler(mall domain.MallRepository) GetStoreHandler {
	return GetStoreHandler{mall: mall}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*domain.MallStore, error) {
	ctx, span := tracer.Start(ctx, "GetStore")
	defer span.End()

	return h.mall.Find(ctx, query.ID)
}
