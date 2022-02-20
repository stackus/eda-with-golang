package application

import (
	"context"
)

type OrderRepository interface {
	Complete(ctx context.Context, orderID string) error
}
