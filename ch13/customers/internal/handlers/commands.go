package handlers

import (
	"context"

	"eda-in-golang/customers/customerspb"
	"eda-in-golang/customers/internal/application"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/registry"
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
	_, err := subscriber.Subscribe(customerspb.CommandChannel, handlers, am.MessageFilter{
		customerspb.AuthorizeCustomerCommand,
	}, am.GroupName("customer-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	ctx, span := tracer.Start(ctx, cmd.CommandName())
	defer span.End()

	switch cmd.CommandName() {
	case customerspb.AuthorizeCustomerCommand:
		return h.doAuthorizeCustomer(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doAuthorizeCustomer(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*customerspb.AuthorizeCustomer)

	return nil, h.app.AuthorizeCustomer(ctx, application.AuthorizeCustomer{ID: payload.GetId()})
}
