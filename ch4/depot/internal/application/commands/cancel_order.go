package commands

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/domain"
)

type CancelOrder struct {
	OrderID string
}

type CancelOrderHandler struct {
	repo domain.ShoppingListRepository
}

func NewCancelOrderHandler(repo domain.ShoppingListRepository) CancelOrderHandler {
	return CancelOrderHandler{repo: repo}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	list, err := h.repo.FindByOrderID(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	err = list.Cancel()
	if err != nil {
		return err
	}

	return h.repo.Update(ctx, list)
}
