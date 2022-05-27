package application

// type CommandHandlers struct{}
//
// func NewCommandHandlers() CommandHandlers {
// 	return CommandHandlers{}
// }
//
// func (h CommandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
// 	switch cmd.CommandName() {
// 	case depotpb.CreateShoppingListCommand:
// 		return h.doCreateShoppingList(ctx, cmd)
// 	case depotpb.CancelShoppingListCommand:
// 		return h.doCancelShoppingList(ctx, cmd)
// 	}
//
// 	return nil, nil
// }
//
// func (h CommandHandlers) doCreateShoppingList(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
// 	payload := cmd.Payload().(*depotpb.CreateShoppingList)
//
// 	list := domain.CreateShopping(cmd.ID(), payload.GetOrderId())
//
// 	for _, item := range payload.GetItems() {
// 		// horribly inefficient
// 		store, err := h.stores.Find(ctx, item.StoreID)
// 		if err != nil {
// 			return errors.Wrap(err, "building shopping list")
// 		}
// 		product, err := h.products.Find(ctx, item.ProductID)
// 		if err != nil {
// 			return errors.Wrap(err, "building shopping list")
// 		}
// 		err = list.AddItem(store, product, item.Quantity)
// 		if err != nil {
// 			return errors.Wrap(err, "building shopping list")
// 		}
// 	}
//
// 	if err := h.shoppingLists.Save(ctx, list); err != nil {
// 		return errors.Wrap(err, "scheduling shopping")
// 	}
//
// 	// publish domain events
// 	if err := h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
// 		return err
// 	}
//
// 	return am.NewReply("TODO", nil, cmd), nil
// }
//
// func (h CommandHandlers) doCancelShoppingList(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
// 	return am.NewReply("TODO", nil, cmd), nil
// }
