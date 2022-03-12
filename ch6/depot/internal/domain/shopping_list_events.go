package domain

const (
	ShoppingListCreatedEvent   = "depot.ShoppingListCreated"
	ShoppingListCanceledEvent  = "depot.ShoppingListCanceled"
	ShoppingListAssignedEvent  = "depot.ShoppingListAssigned"
	ShoppingListCompletedEvent = "depot.ShoppingListCompleted"
)

type ShoppingListCreated struct {
	ShoppingList *ShoppingList
}

type ShoppingListCanceled struct {
	ShoppingList *ShoppingList
}

type ShoppingListAssigned struct {
	ShoppingList *ShoppingList
	BotID        string
}

type ShoppingListCompleted struct {
	ShoppingList *ShoppingList
}
