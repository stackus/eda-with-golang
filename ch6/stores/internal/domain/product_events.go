package domain

const (
	ProductAddedEvent   = "stores.ProductAdded"
	ProductRemovedEvent = "stores.ProductRemoved"
)

type ProductAdded struct {
	Product *Product
}

type ProductRemoved struct {
	Product *Product
}
