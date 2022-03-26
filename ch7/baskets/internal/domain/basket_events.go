package domain

const (
	BasketStartedEvent     = "baskets.BasketStarted"
	BasketItemAddedEvent   = "baskets.BasketItemAdded"
	BasketItemRemovedEvent = "baskets.BasketItemRemoved"
	BasketCanceledEvent    = "baskets.BasketCanceled"
	BasketCheckedOutEvent  = "baskets.BasketCheckedOut"
)

type BasketStarted struct {
	Basket *Basket
}

// Key implements registry.Registerable
func (BasketStarted) Key() string { return BasketStartedEvent }

type BasketItemAdded struct {
	Basket *Basket
	Item   Item
}

// Key implements registry.Registerable
func (BasketItemAdded) Key() string { return BasketItemAddedEvent }

type BasketItemRemoved struct {
	Basket *Basket
	Item   Item
}

// Key implements registry.Registerable
func (BasketItemRemoved) Key() string { return BasketItemRemovedEvent }

type BasketCanceled struct {
	Basket *Basket
}

// Key implements registry.Registerable
func (BasketCanceled) Key() string { return BasketCanceledEvent }

type BasketCheckedOut struct {
	Basket *Basket
}

// Key implements registry.Registerable
func (BasketCheckedOut) Key() string { return BasketCheckedOutEvent }
