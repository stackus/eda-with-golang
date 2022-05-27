package orderingpb

func (*OrderCreated) Key() string   { return OrderCreatedEvent }
func (*OrderReadied) Key() string   { return OrderReadiedEvent }
func (*OrderCanceled) Key() string  { return OrderCanceledEvent }
func (*OrderCompleted) Key() string { return OrderCompletedEvent }

func (*RejectOrder) Key() string  { return RejectOrderCommand }
func (*ApproveOrder) Key() string { return ApproveOrderCommand }
