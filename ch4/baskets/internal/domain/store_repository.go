package domain

import (
	"context"
)

type StoreRepository interface {
	Find(ctx context.Context, id StoreID) (*Store, error)
}
