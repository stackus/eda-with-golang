package handlers

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/em"
	"eda-in-golang/ch7/stores/storespb"
)

func RegisterProductHandlers(productHandlers ddd.EventHandler[ddd.Event], stream em.EventSubscriber) error {
	evtMsgHandler := em.MessageHandlerFunc[em.EventMessage](func(ctx context.Context, eventMsg em.EventMessage) error {
		return productHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.ProductAggregateChannel, evtMsgHandler, em.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductPriceIncreasedEvent,
		storespb.ProductPriceDecreasedEvent,
		storespb.ProductRemovedEvent,
	})
}
