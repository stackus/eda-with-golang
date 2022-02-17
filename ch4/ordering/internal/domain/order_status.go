package domain

type OrderStatus string

const (
	OrderUnknown   OrderStatus = ""
	OrderPending   OrderStatus = "pending"
	OrderInProcess OrderStatus = "in-progress"
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

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case OrderPending.String():
		return OrderPending
	case OrderInProcess.String():
		return OrderInProcess
	case OrderCancelled.String():
		return OrderCancelled
	case OrderCompleted.String():
		return OrderCompleted
	default:
		return OrderUnknown
	}
}
