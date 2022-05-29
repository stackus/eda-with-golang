package cosec

import (
	"context"

	"eda-in-golang/ch9/cosec/internal"
	"eda-in-golang/ch9/cosec/internal/handlers"
	"eda-in-golang/ch9/cosec/internal/logging"
	"eda-in-golang/ch9/cosec/internal/models"
	"eda-in-golang/ch9/customers/customerspb"
	"eda-in-golang/ch9/depot/depotpb"
	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
	"eda-in-golang/ch9/internal/jetstream"
	"eda-in-golang/ch9/internal/monolith"
	pg "eda-in-golang/ch9/internal/postgres"
	"eda-in-golang/ch9/internal/registry"
	"eda-in-golang/ch9/internal/registry/serdes"
	"eda-in-golang/ch9/internal/sec"
	"eda-in-golang/ch9/ordering/orderingpb"
	"eda-in-golang/ch9/payments/paymentspb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = depotpb.Registrations(reg); err != nil {
		return err
	}
	if err = paymentspb.Registrations(reg); err != nil {
		return err
	}
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	sagaStore := pg.NewSagaStore("cosec.sagas", mono.DB(), reg)
	createOrderSagaRepo := sec.NewSagaRepository[*models.CreateOrderData](reg, sagaStore)

	// setup application
	createOrderSaga := logging.LogReplyHandlerAccess[*models.CreateOrderData](
		sec.NewOrchestrator[*models.CreateOrderData](internal.NewCreateOrderSaga(), createOrderSagaRepo, commandStream),
		"CreateOrderSaga", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(createOrderSaga),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterReplies(replyStream, createOrderSaga); err != nil {
		return err
	}

	return
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}
