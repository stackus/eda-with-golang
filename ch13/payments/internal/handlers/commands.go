package handlers

import (
	"context"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/payments/internal/application"
	"eda-in-golang/payments/paymentspb"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(app application.App) ddd.CommandHandler[ddd.Command] {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(paymentspb.CommandChannel, handlers, am.MessageFilter{
		paymentspb.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	ctx, span := tracer.Start(ctx, cmd.CommandName())
	defer span.End()

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
