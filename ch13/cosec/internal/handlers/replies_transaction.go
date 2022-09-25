package handlers

import (
	"context"
	"database/sql"

	"eda-in-golang/cosec/internal/models"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/di"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/sec"
	"eda-in-golang/internal/tm"
)

func RegisterReplyHandlersTx(container di.Container) error {
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

		return am.NewReplyHandler(
			di.Get(ctx, "registry").(registry.Registry),
			di.Get(ctx, "orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
			tm.InboxHandler(di.Get(ctx, "inboxStore").(tm.InboxStore)),
		).HandleMessage(ctx, msg)
	})

	subscriber := container.Get("messageSubscriber").(am.MessageSubscriber)

	return RegisterReplyHandlers(subscriber, rawMsgHandler)
}
