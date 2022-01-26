package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type (
	CreateStore struct {
		Name     string
		Location string
	}

	CreateStoreHandler struct {
		repo ports.StoreRepository
	}
)

func NewCreateStoreHandler(repo ports.StoreRepository) CreateStoreHandler {
	return CreateStoreHandler{repo: repo}
}

func (h CreateStoreHandler) Handle(ctx context.Context, cmd CreateStore) (string, error) {
	storeID := uuid.New().String()
	store, err := domain.CreateStore(storeID, cmd.Name, domain.NewLocation(cmd.Location))
	if err != nil {
		return "", err
	}

	err = h.repo.SaveStore(ctx, store)

	return storeID, err
}
