package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
)

type DisableParticipation struct {
	ID string
}

type DisableParticipationHandler struct {
	repo ports.StoreRepository
}

func NewDisableParticipationHandler(repo ports.StoreRepository) DisableParticipationHandler {
	return DisableParticipationHandler{repo: repo}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.repo.FindStore(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.DisableParticipation()
	if err != nil {
		return err
	}

	return h.repo.UpdateStore(ctx, store)
}
