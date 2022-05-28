package handlers

import (
	"context"

	"eda-in-golang/ch9/internal/ac"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/ordering/internal/domain"
)

func RegisterCreateOrderReplies(subscriber am.ReplySubscriber, saga ac.Orchestrator[*domain.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return saga.HandleReply(ctx, replyMsg)
	})
	return subscriber.Subscribe(saga.ReplyTopic(), replyMsgHandler, am.GroupName("ordering-create-replies"))
}
