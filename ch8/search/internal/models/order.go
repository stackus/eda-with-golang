package models

type Order struct {
	OrderID      string
	CustomerID   string
	CustomerName string
	Items        []Item
	Total        float64
	Status       string
}

type Item struct {
	ProductID   string
	StoreID     string
	ProductName string
	StoreName   string
	Price       float64
	Quantity    int64
}
