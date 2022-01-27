package queries

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type GetOffering struct {
	ID      string
	StoreID string
}

type GetOfferingHandler struct {
	repo ports.OfferingRepository
}

func NewGetOfferingHandler(repo ports.OfferingRepository) GetOfferingHandler {
	return GetOfferingHandler{repo: repo}
}

func (h GetOfferingHandler) GetOffering(ctx context.Context, query GetOffering) (*domain.Offering, error) {
	return h.repo.FindOffering(ctx, query.ID, query.StoreID)
}
