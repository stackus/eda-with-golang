package handlers

// func RegisterCommandHandlers(cmdHandlers am.CommandHandler[am.Command], stream am.CommandStream) error {
// 	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (am.Reply, error) {
// 		return cmdHandlers.HandleCommand(ctx, cmdMsg)
// 	})
//
// 	return stream.Subscribe(depotpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
// 		depotpb.CreateShoppingListCommand,
// 		depotpb.CancelShoppingListCommand,
// 	}, am.GroupName("depot-commands"))
// }
