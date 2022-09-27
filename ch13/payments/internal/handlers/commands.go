package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/registry"
	"eda-in-golang/payments/internal/application"
	"eda-in-golang/payments/paymentspb"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(reg registry.Registry, app application.App, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app: app,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(paymentspb.CommandChannel, handlers, am.MessageFilter{
		paymentspb.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		attrs := []attribute.KeyValue{
			attribute.String("Command", cmd.CommandName()),
			attribute.Float64("Took", time.Since(started).Seconds()),
		}
		if err != nil {
			attrs = append(attrs, attribute.String("Error", err.Error()))
		}
		span.AddEvent("Handled Command", trace.WithAttributes(attrs...))
	}(time.Now())

	switch cmd.CommandName() {
	case paymentspb.ConfirmPaymentCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doConfirmPayment(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*paymentspb.ConfirmPayment)

	return nil, h.app.ConfirmPayment(ctx, application.ConfirmPayment{ID: payload.GetId()})
}
