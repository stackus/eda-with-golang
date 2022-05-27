package handlers

import (
	"eda-in-golang/ch9/internal/ac"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/ordering/internal/domain"
)

func RegisterCreateOrderReplies(subscriber am.ReplySubscriber, saga ac.Orchestrator[*domain.CreateOrderData]) error {
	return subscriber.Subscribe(saga.ReplyTopic(), saga, am.GroupName("ordering-create-replies"))
}
