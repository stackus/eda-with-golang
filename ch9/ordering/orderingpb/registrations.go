package orderingpb

import (
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/internal/registry/serdes"
)

const (
	OrderAggregateChannel = "mallbots.ordering.events.Order"

	OrderCreatedEvent   = "ordersapi.OrderCreated"
	OrderReadiedEvent   = "ordersapi.OrderReadied"
	OrderCanceledEvent  = "ordersapi.OrderCanceled"
	OrderCompletedEvent = "ordersapi.OrderCompleted"

	CommandChannel          = "mallbots.ordering.commands"
	CreateOrderReplyChannel = "mallbots.ordering.replies.CreateOrder"

	RejectOrderCommand  = "ordersapi.RejectOrder"
	ApproveOrderCommand = "ordersapi.ApproveOrder"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Order events
	if err = serde.Register(&OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderCompleted{}); err != nil {
		return err
	}

	if err = serde.Register(&RejectOrder{}); err != nil {
		return err
	}
	if err = serde.Register(&ApproveOrder{}); err != nil {
		return err
	}

	return nil
}
