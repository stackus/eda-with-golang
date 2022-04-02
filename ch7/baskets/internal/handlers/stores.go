package handlers

import (
	"context"

	"eda-in-golang/ch7/internal/am"
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/stores/storespb"
)

func RegisterStoreHandlers(storeHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreParticipatingToggledEvent,
		storespb.StoreRebrandedEvent,
	})
}
