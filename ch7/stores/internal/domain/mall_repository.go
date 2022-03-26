package domain

import (
	"context"
)

type MallRepository interface {
	AddStore(ctx context.Context, storeID, name, location string) error
	SetStoreParticipation(ctx context.Context, storeID string, participating bool) error
	RenameStore(ctx context.Context, storeID, name string) error
	Find(ctx context.Context, storeID string) (*Store, error)
	All(ctx context.Context) ([]*Store, error)
	AllParticipating(ctx context.Context) ([]*Store, error)
}
