package domain

type ProductAdded struct {
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func (ProductAdded) EventName() string { return "stores.ProductAdded" }

type ProductRemoved struct{}

func (ProductRemoved) EventName() string { return "stores.ProductRemoved" }
