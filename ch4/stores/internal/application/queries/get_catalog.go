package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetCatalog struct {
	StoreID string
}

type GetCatalogHandler struct {
	repo domain.ProductRepository
}

func NewGetCatalogHandler(repo domain.ProductRepository) GetCatalogHandler {
	return GetCatalogHandler{repo: repo}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.Product, error) {
	return h.repo.GetCatalog(ctx, query.StoreID)
}
