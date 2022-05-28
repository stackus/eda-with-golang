package domain

type CreateOrderData struct {
	OrderID    string
	CustomerID string
	PaymentID  string
	ShoppingID string
	Items      []Item
	Total      float64
}
