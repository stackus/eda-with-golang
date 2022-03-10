package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch5/internal/ddd"
)

type DomainEventHandlers interface {
	OnShoppingListCreated(ctx context.Context, event ddd.Event) error
	OnShoppingListCanceled(ctx context.Context, event ddd.Event) error
	OnShoppingListAssigned(ctx context.Context, event ddd.Event) error
	OnShoppingListCompleted(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnShoppingListCreated(ctx context.Context, event ddd.Event) error {
	// TODO implement me
	panic("implement me")
}

func (ignoreUnimplementedDomainEvents) OnShoppingListCanceled(ctx context.Context, event ddd.Event) error {
	// TODO implement me
	panic("implement me")
}

func (ignoreUnimplementedDomainEvents) OnShoppingListAssigned(ctx context.Context, event ddd.Event) error {
	// TODO implement me
	panic("implement me")
}

func (ignoreUnimplementedDomainEvents) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	// TODO implement me
	panic("implement me")
}
