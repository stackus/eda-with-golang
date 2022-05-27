package depotpb

// Commands
func (*CreateShoppingList) Key() string { return CreateShoppingListCommand }
func (*CancelShoppingList) Key() string { return CancelShoppingListCommand }

// Replies
func (*CreatedShoppingList) Key() string { return CreatedShoppingListReply }
