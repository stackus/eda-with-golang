package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetProduct struct {
	ID string
}

type GetProductHandler struct {
	repo domain.ProductRepository
}

func NewGetProductHandler(repo domain.ProductRepository) GetProductHandler {
	return GetProductHandler{repo: repo}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.Product, error) {
	return h.repo.FindProduct(ctx, query.ID)
}
