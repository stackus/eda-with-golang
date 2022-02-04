package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type (
	CreateStore struct {
		ID       string
		Name     string
		Location string
	}

	CreateStoreHandler struct {
		repo domain.StoreRepository
	}
)

func NewCreateStoreHandler(repo domain.StoreRepository) CreateStoreHandler {
	return CreateStoreHandler{repo: repo}
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {
	store, err := domain.CreateStore(cmd.ID, cmd.Name, domain.NewLocation(cmd.Location))
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, store)

	return err
}
