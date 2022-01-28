package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type AddOffering struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	Price       float64
}

type AddOfferingHandler struct {
	storeRepo    ports.StoreRepository
	offeringRepo ports.OfferingRepository
}

func NewAddOfferingHandler(storeRepo ports.StoreRepository, offeringRepo ports.OfferingRepository) AddOfferingHandler {
	return AddOfferingHandler{
		storeRepo:    storeRepo,
		offeringRepo: offeringRepo,
	}
}

func (h AddOfferingHandler) AddOffering(ctx context.Context, cmd AddOffering) error {
	_, err := h.storeRepo.Find(ctx, cmd.StoreID)
	if err != nil {
		return errors.Wrap(err, "error adding offering")
	}

	offering, err := domain.CreateOffering(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "error adding offering")
	}

	return errors.Wrap(h.offeringRepo.AddOffering(ctx, offering), "error adding offering")
}
