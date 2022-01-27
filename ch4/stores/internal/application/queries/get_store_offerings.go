package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetStoreOfferings struct {
	StoreID string
}

type GetStoreOfferingsHandler struct {
	repo ports.OfferingRepository
}

func NewGetStoreOfferingsHandler(repo ports.OfferingRepository) GetStoreOfferingsHandler {
	return GetStoreOfferingsHandler{repo: repo}
}

func (h GetStoreOfferingsHandler) GetStoreOfferings(ctx context.Context, query GetStoreOfferings) ([]*domain.Offering, error) {
	return h.repo.GetStoreOfferings(ctx, query.StoreID)
}
