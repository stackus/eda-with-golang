package handlers

import (
	"context"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/ordering/orderingpb"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return orderHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(orderingpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderingpb.OrderCreatedEvent,
		orderingpb.OrderReadiedEvent,
		orderingpb.OrderCanceledEvent,
		orderingpb.OrderCompletedEvent,
	}, am.GroupName("notification-orders"))
}
