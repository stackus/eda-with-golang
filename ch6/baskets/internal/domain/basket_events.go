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

type BasketItemAdded struct {
	Basket *Basket
	Item   Item
}

type BasketItemRemoved struct {
	Basket *Basket
	Item   Item
}

type BasketCanceled struct {
	Basket *Basket
}

type BasketCheckedOut struct {
	Basket *Basket
}
