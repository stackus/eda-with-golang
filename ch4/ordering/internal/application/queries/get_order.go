package queries

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	repo ports.OrderRepository
}

func NewGetOrderHandler(repo ports.OrderRepository) GetOrderHandler {
	return GetOrderHandler{repo: repo}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.repo.Find(ctx, query.ID)

	return order, errors.Wrap(err, "get order query")
}
