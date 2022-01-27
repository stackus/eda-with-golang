package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
)

type RemoveOffering struct {
	ID      string
	StoreID string
}

type RemoveOfferingHandler struct {
	storeRepo    ports.StoreRepository
	offeringRepo ports.OfferingRepository
}

func NewRemoveOfferingHandler(storeRepo ports.StoreRepository, offeringRepo ports.OfferingRepository) RemoveOfferingHandler {
	return RemoveOfferingHandler{
		storeRepo:    storeRepo,
		offeringRepo: offeringRepo,
	}
}

func (h RemoveOfferingHandler) RemoveOffering(ctx context.Context, cmd RemoveOffering) error {
	_, err := h.storeRepo.FindStore(ctx, cmd.StoreID)
	if err != nil {
		return err
	}

	return h.offeringRepo.RemoveOffering(ctx, cmd.ID, cmd.StoreID)
}
