package handlers

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/em"
	"eda-in-golang/ch7/stores/storespb"
)

func RegisterStoreHandlers(storeHandlers ddd.EventHandler[ddd.Event], stream em.EventSubscriber) error {
	evtMsgHandler := em.MessageHandlerFunc[em.EventMessage](func(ctx context.Context, eventMsg em.EventMessage) error {
		return storeHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, em.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreParticipatingToggledEvent,
		storespb.StoreRebrandedEvent,
	})
}
