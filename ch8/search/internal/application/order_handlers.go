package application

import (
	"context"

	"eda-in-golang/ch8/internal/ddd"
	"eda-in-golang/ch8/ordering/orderingpb"
)

type OrderHandlers[T ddd.Event] struct {
	orders    OrderRepository
	customers CustomerRepository
	stores    StoreRepository
	products  ProductRepository
}

var _ ddd.EventHandler[ddd.Event] = (*OrderHandlers[ddd.Event])(nil)

func NewOrderHandlers(orders OrderRepository, customers CustomerRepository, stores StoreRepository, products ProductRepository) OrderHandlers[ddd.Event] {
	return OrderHandlers[ddd.Event]{
		orders:    orders,
		customers: customers,
		stores:    stores,
		products:  products,
	}
}

func (h OrderHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderingpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderingpb.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderingpb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case orderingpb.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	}
	return nil
}

func (h OrderHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCreated)
	// TODO implement me
	panic("implement me")
}

func (h OrderHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderReadied)
	// TODO implement me
	panic("implement me")
}

func (h OrderHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	// TODO implement me
	panic("implement me")
}

func (h OrderHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCompleted)
	// TODO implement me
	panic("implement me")
}
