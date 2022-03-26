package domain

const (
	OrderCreatedEvent   = "ordering.OrderCreated"
	OrderCanceledEvent  = "ordering.OrderCanceled"
	OrderReadiedEvent   = "ordering.OrderReadied"
	OrderCompletedEvent = "ordering.OrderCompleted"
)

type OrderCreated struct {
	Order *Order
}

func (OrderCreated) Key() string { return OrderCreatedEvent }

type OrderCanceled struct {
	Order *Order
}

func (OrderCanceled) Key() string { return OrderCanceledEvent }

type OrderReadied struct {
	Order *Order
}

func (OrderReadied) Key() string { return OrderReadiedEvent }

type OrderCompleted struct {
	Order *Order
}

func (OrderCompleted) Key() string { return OrderCompletedEvent }
