package application

import (
	"context"

	"eda-in-golang/ch9/depot/depotpb"
	"eda-in-golang/ch9/internal/ac"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/ordering/internal/domain"
	"eda-in-golang/ch9/ordering/orderingpb"
)

const CreateOrderSagaName = "ordering.CreateOrder"

type createOrderSaga struct {
	ac.Saga[*domain.CreateOrderData]
}

func NewCreateOrderSaga() ac.Saga[*domain.CreateOrderData] {
	saga := createOrderSaga{
		Saga: ac.NewSaga[*domain.CreateOrderData](CreateOrderSagaName, orderingpb.CreateOrderReplyChannel),
	}

	// 0. -RejectOrder
	saga.AddStep().
		Compensation(saga.rejectOrder)

	// 1. AuthorizeCustomer
	saga.AddStep().
		Action(saga.authorizeCustomer)

	// 2. CreateShoppingList, -CancelShoppingList
	saga.AddStep().
		Action(saga.createShoppingList).
		OnActionReply(depotpb.CreatedShoppingListReply, saga.onCreateShoppingListReply).
		Compensation(saga.cancelShoppingList)

	// 3. ConfirmPayment
	saga.AddStep().
		Action(saga.confirmPayment)

	// 4. InitiateShopping
	saga.AddStep().
		Action(saga.initiateShopping)

	// 5. ApproveOrder
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return am.NewCommand(orderingpb.RejectOrderCommand, orderingpb.CommandChannel, &orderingpb.RejectOrder{Id: data.OrderID})
}

func (s createOrderSaga) authorizeCustomer(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return nil
}

func (s createOrderSaga) createShoppingList(ctx context.Context, data *domain.CreateOrderData) am.Command {
	items := make([]*depotpb.CreateShoppingList_Item, len(data.Items))
	for i, item := range data.Items {
		items[i] = &depotpb.CreateShoppingList_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		}
	}

	return am.NewCommand(depotpb.CreateShoppingListCommand, depotpb.CommandChannel, &depotpb.CreateShoppingList{
		OrderId: data.OrderID,
		Items:   items,
	})
}

func (s createOrderSaga) onCreateShoppingListReply(ctx context.Context, data *domain.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*depotpb.CreatedShoppingList)

	data.ShoppingID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelShoppingList(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return am.NewCommand(depotpb.CancelShoppingListCommand, depotpb.CommandChannel, &depotpb.CancelShoppingList{Id: data.ShoppingID})
}

func (s createOrderSaga) confirmPayment(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return nil
}

func (s createOrderSaga) initiateShopping(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return nil
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return am.NewCommand(orderingpb.ApproveOrderCommand, orderingpb.CommandChannel, &orderingpb.ApproveOrder{Id: data.OrderID})
}
