package domain

type ProductID string
type Price float64

type Product struct {
	ID      ProductID
	StoreID StoreID
	Name    string
	Price   Price
}

func (i ProductID) String() string {
	return string(i)
}

func (p Price) Float64() float64 {
	return float64(p)
}
