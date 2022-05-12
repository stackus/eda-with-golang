package handlers

import (
	"context"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/stores/storespb"
)

func RegisterStoreHandlers(storeHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreRebrandedEvent,
	}, am.GroupName("baskets-stores"))
}
