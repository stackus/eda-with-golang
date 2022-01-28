package domain

type OrderStatus string

const (
	OrderUnknown   OrderStatus = ""
	OrderPending   OrderStatus = "pending"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case OrderPending, OrderCompleted, OrderCancelled:
		return string(s)
	default:
		return ""
	}
}
