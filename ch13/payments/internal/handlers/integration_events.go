package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/registry"
	"eda-in-golang/ordering/orderingpb"
	"eda-in-golang/payments/internal/application"
)

type integrationHandlers[T ddd.Event] struct {
	app application.App
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, app application.App, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		app: app,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(orderingpb.OrderAggregateChannel, handlers, am.MessageFilter{
		orderingpb.OrderReadiedEvent,
	}, am.GroupName("payment-orders"))
	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		attrs := []attribute.KeyValue{
			attribute.String("Event", event.EventName()),
			attribute.Float64("Took", time.Since(started).Seconds()),
		}
		if err != nil {
			attrs = append(attrs, attribute.String("Error", err.Error()))
		}
		span.AddEvent("Handled Integration Event", trace.WithAttributes(attrs...))
	}(time.Now())

	switch event.EventName() {
	case orderingpb.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderingpb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}
	return nil
}

func (h integrationHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderReadied)
	return h.app.CreateInvoice(ctx, application.CreateInvoice{
		ID:        payload.GetId(),
		OrderID:   payload.GetId(),
		PaymentID: payload.GetPaymentId(),
		Amount:    payload.GetTotal(),
	})
}

func (h integrationHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	return h.app.CancelInvoice(ctx, application.CancelInvoice{
		ID: payload.GetId(),
	})
}
