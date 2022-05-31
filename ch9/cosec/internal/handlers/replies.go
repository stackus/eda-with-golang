package handlers

import (
	"context"

	"eda-in-golang/cosec/internal/models"
	"eda-in-golang/internal/am"
	"eda-in-golang/internal/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscriber, saga sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return saga.HandleReply(ctx, replyMsg)
	})
	return subscriber.Subscribe(saga.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
}
