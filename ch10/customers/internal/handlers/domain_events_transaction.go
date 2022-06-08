package handlers

import (
	"context"

	"eda-in-golang/customers/internal/domain"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
)

func RegisterDomainEventHandlersTx(subscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		domainHandlers := di.Get(ctx, "domainEventHandlers").(ddd.EventHandler[ddd.AggregateEvent])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber.Subscribe(handlers,
		domain.CustomerRegisteredEvent,
		domain.CustomerSmsChangedEvent,
		domain.CustomerEnabledEvent,
		domain.CustomerDisabledEvent,
	)
}
