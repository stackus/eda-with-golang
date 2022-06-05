package handlers

import (
	"context"

	"eda-in-golang/depot/depotpb"
	"eda-in-golang/depot/internal/domain"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) ddd.EventHandler[ddd.AggregateEvent] {
	return domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handler ddd.EventHandler[ddd.AggregateEvent]) {
	/*

		handler := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event T) error {
			domainHandler := di.Get(ctx, "domainHandler").(ddd.EventHandler[ddd.AggregateEvent])

			return domainHandler.HandleEvent(ctx, event)
		})

	*/

	subscriber.Subscribe(handler, domain.ShoppingListCompletedEvent)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onShoppingListCompleted(ctx context.Context, event ddd.AggregateEvent) error {
	completed := event.Payload().(*domain.ShoppingListCompleted)

	return h.publisher.Publish(ctx, depotpb.ShoppingListAggregateChannel, ddd.NewEvent(depotpb.ShoppingListCompletedEvent, &depotpb.ShoppingListCompleted{
		Id:      event.AggregateID(),
		OrderId: completed.ShoppingList.OrderID,
	}))
}
