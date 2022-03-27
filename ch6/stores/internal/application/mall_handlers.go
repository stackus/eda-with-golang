package application

import (
	"context"

	"eda-in-golang/ch6/internal/ddd"
	"eda-in-golang/ch6/stores/internal/domain"
)

type MallHandlers struct {
	mall domain.MallRepository
}

var _ ddd.EventHandler = (*MallHandlers)(nil)

func NewMallHandlers(mall domain.MallRepository) *MallHandlers {
	return &MallHandlers{
		mall: mall,
	}
}

func (h MallHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domain.StoreParticipationEnabledEvent:
		return h.onStoreParticipationEnabled(ctx, event)
	case domain.StoreParticipationDisabledEvent:
		return h.onStoreParticipationDisabled(ctx, event)
	case domain.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}
	return nil
}

func (h MallHandlers) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.mall.AddStore(ctx, event.AggregateID(), payload.Name, payload.Location)
}

func (h MallHandlers) onStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), true)
}

func (h MallHandlers) onStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), false)
}

func (h MallHandlers) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreRebranded)
	return h.mall.RenameStore(ctx, event.AggregateID(), payload.Name)
}
