package handlers

import (
	"context"

	"eda-in-golang/ch9/depot/depotpb"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/payments/internal/application"
	"eda-in-golang/ch9/payments/paymentspb"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(app application.App) am.CommandHandler {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.CommandSubscriber, handlers am.CommandHandler) error {
	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (ddd.Reply, error) {
		return handlers.HandleCommand(ctx, cmdMsg)
	})

	return subscriber.Subscribe(depotpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		paymentspb.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case depotpb.CreateShoppingListCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doConfirmPayment(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*paymentspb.ConfirmPayment)

	return nil, h.app.ConfirmPayment(ctx, application.ConfirmPayment{ID: payload.GetId()})
}
