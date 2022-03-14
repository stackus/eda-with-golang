package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type DomainEventHandlers interface {
	OnStoreCreated(ctx context.Context, event ddd.Event) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error
	OnStoreRebranded(ctx context.Context, event ddd.Event) error

	OnProductAdded(ctx context.Context, event ddd.Event) error
	OnProductRebranded(ctx context.Context, event ddd.Event) error
	OnProductPriceIncreased(ctx context.Context, event ddd.Event) error
	OnProductPriceDecreased(ctx context.Context, event ddd.Event) error
	OnProductRemoved(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnStoreCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreRebranded(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductAdded(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductRebranded(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductPriceIncreased(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductPriceDecreased(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductRemoved(ctx context.Context, event ddd.Event) error {
	return nil
}
