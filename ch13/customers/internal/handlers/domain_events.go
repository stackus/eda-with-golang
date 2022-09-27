package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"eda-in-golang/customers/customerspb"
	"eda-in-golang/customers/internal/domain"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domain.CustomerRegisteredEvent,
		domain.CustomerSmsChangedEvent,
		domain.CustomerEnabledEvent,
		domain.CustomerDisabledEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		attrs := []attribute.KeyValue{
			attribute.String("Event", event.EventName()),
			attribute.Float64("Took", time.Since(started).Seconds()),
		}
		if err != nil {
			attrs = append(attrs, attribute.String("Error", err.Error()))
		}
		span.AddEvent("Handled Domain Event", trace.WithAttributes(attrs...))
	}(time.Now())

	switch event.EventName() {
	case domain.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	case domain.CustomerSmsChangedEvent:
		return h.onCustomerSmsChanged(ctx, event)
	case domain.CustomerEnabledEvent:
		return h.onCustomerEnabled(ctx, event)
	case domain.CustomerDisabledEvent:
		return h.onCustomerDisabled(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onCustomerRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.CustomerRegistered)
	return h.publisher.Publish(ctx, customerspb.CustomerAggregateChannel,
		ddd.NewEvent(customerspb.CustomerRegisteredEvent, &customerspb.CustomerRegistered{
			Id:        payload.Customer.ID(),
			Name:      payload.Customer.Name,
			SmsNumber: payload.Customer.SmsNumber,
		}),
	)
}

func (h domainHandlers[T]) onCustomerSmsChanged(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.CustomerRegistered)
	return h.publisher.Publish(ctx, customerspb.CustomerAggregateChannel,
		ddd.NewEvent(customerspb.CustomerSmsChangedEvent, &customerspb.CustomerSmsChanged{
			Id:        payload.Customer.ID(),
			SmsNumber: payload.Customer.SmsNumber,
		}),
	)
}

func (h domainHandlers[T]) onCustomerEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerspb.CustomerAggregateChannel,
		ddd.NewEvent(customerspb.CustomerEnabledEvent, &customerspb.CustomerEnabled{
			Id: event.AggregateID(),
		}),
	)
}

func (h domainHandlers[T]) onCustomerDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerspb.CustomerAggregateChannel,
		ddd.NewEvent(customerspb.CustomerDisabledEvent, &customerspb.CustomerDisabled{
			Id: event.AggregateID(),
		}),
	)
}
