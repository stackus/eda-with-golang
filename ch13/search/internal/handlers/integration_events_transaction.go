package handlers

import (
	"context"
	"database/sql"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/tm"
)

func RegisterIntegrationEventHandlersTx(container di.Container) (err error) {
	rawMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
		ctx = container.Scoped(ctx)
		defer func(tx *sql.Tx) {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}(di.Get(ctx, "tx").(*sql.Tx))

		return am.NewEventHandler(
			di.Get(ctx, "registry").(registry.Registry),
			di.Get(ctx, "integrationEventHandlers").(ddd.EventHandler[ddd.Event]),
			tm.InboxHandler(di.Get(ctx, "inboxStore").(tm.InboxStore)),
		).HandleMessage(ctx, msg)
	})

	subscriber := container.Get("messageSubscriber").(am.MessageSubscriber)

	return RegisterIntegrationEventHandlers(subscriber, rawMsgHandler)
}