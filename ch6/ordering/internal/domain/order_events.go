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

type OrderCanceled struct {
	Order *Order
}

type OrderReadied struct {
	Order *Order
}

type OrderCompleted struct {
	Order *Order
}
