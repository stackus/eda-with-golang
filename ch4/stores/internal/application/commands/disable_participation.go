package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type DisableParticipation struct {
	ID string
}

type DisableParticipationHandler struct {
	repo domain.StoreRepository
}

func NewDisableParticipationHandler(repo domain.StoreRepository) DisableParticipationHandler {
	return DisableParticipationHandler{repo: repo}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.repo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.DisableParticipation()
	if err != nil {
		return err
	}

	return h.repo.Update(ctx, store)
}
