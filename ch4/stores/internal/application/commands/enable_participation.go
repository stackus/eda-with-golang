package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	repo ports.StoreRepository
}

func NewEnableParticipationHandler(repo ports.StoreRepository) EnableParticipationHandler {
	return EnableParticipationHandler{repo: repo}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.repo.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.EnableParticipation()
	if err != nil {
		return err
	}

	return h.repo.Update(ctx, store)
}
