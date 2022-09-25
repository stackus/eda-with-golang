package handlers

import (
	"eda-in-golang/cosec/internal"
	"eda-in-golang/internal/am"
)

func RegisterReplyHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(internal.CreateOrderReplyChannel, handlers, am.GroupName("cosec-replies"))
	return err
}
