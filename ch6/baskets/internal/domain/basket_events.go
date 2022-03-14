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

func (BasketStarted) Key() string { return BasketStartedEvent }

type BasketItemAdded struct {
	Basket *Basket
	Item   Item
}

func (BasketItemAdded) Key() string { return BasketItemAddedEvent }

type BasketItemRemoved struct {
	Basket *Basket
	Item   Item
}

func (BasketItemRemoved) Key() string { return BasketItemRemovedEvent }

type BasketCanceled struct {
	Basket *Basket
}

func (BasketCanceled) Key() string { return BasketCanceledEvent }

type BasketCheckedOut struct {
	Basket *Basket
}

func (BasketCheckedOut) Key() string { return BasketCheckedOutEvent }
