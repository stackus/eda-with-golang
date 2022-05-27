package depotpb

import (
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/internal/registry/serdes"
)

const (
	CommandChannel = "mallbots.depot.commands"

	CreateShoppingListCommand = "depotapi.CreateShoppingListCommand"
	CancelShoppingListCommand = "depotapi.CancelShoppingListCommand"

	CreatedShoppingListReply = "depotapi.CreatedShoppingListReply"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	if err = serde.Register(&CreateShoppingList{}); err != nil {
		return err
	}
	if err = serde.Register(&CancelShoppingList{}); err != nil {
		return err
	}

	if err = serde.Register(&CreatedShoppingList{}); err != nil {
		return err
	}

	return nil
}
