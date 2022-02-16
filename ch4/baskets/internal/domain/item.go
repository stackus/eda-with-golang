package domain

type Item struct {
	StoreID      StoreID
	ProductID    ProductID
	StoreName    string
	ProductName  string
	ProductPrice Price
	Quantity     int
}
