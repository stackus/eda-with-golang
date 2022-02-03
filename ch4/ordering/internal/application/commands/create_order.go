package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type CreateOrder struct {
	ID        string
	Items     []*domain.Item
	CardToken string
	SmsNumber string
}

type CreateOrderHandler struct {
	repo domain.OrderRepository
}

func NewCreateOrderHandler(repo domain.OrderRepository) CreateOrderHandler {
	return CreateOrderHandler{repo: repo}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.Items, cmd.CardToken, cmd.SmsNumber)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	return errors.Wrap(h.repo.Save(ctx, order), "create order command")
}
