package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type GetCatalog struct {
	StoreID string
}

type GetCatalogHandler struct {
	catalog domain.CatalogRepository
}

func NewGetCatalogHandler(catalog domain.CatalogRepository) GetCatalogHandler {
	return GetCatalogHandler{catalog: catalog}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.Product, error) {
	return h.catalog.GetCatalog(ctx, query.StoreID)
}
