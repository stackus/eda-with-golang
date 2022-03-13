package application

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

type MallHandlers struct {
	mall domain.MallRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*MallHandlers)(nil)

func NewMallHandlers(mall domain.MallRepository) *MallHandlers {
	return &MallHandlers{
		mall: mall,
	}
}

func (h MallHandlers) OnStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.mall.AddStore(ctx, event.AggregateID(), payload.Name, payload.Location)
}

func (h MallHandlers) OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), true)
}

func (h MallHandlers) OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), false)
}

func (h MallHandlers) OnStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreRebranded)
	return h.mall.RenameStore(ctx, event.AggregateID(), payload.Name)
}
