package ports

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/domain"
)

type OrderRepository interface {
	Find(ctx context.Context, orderID string) (*domain.Order, error)
	Save(ctx context.Context, order *domain.Order) error
	Update(ctx context.Context, order *domain.Order) error
}
