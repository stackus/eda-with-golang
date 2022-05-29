package handlers

import (
	"context"

	"eda-in-golang/ch9/cosec/internal/models"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/sec"
)

func RegisterReplies(subscriber am.ReplySubscriber, saga sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return saga.HandleReply(ctx, replyMsg)
	})
	return subscriber.Subscribe(saga.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
}
