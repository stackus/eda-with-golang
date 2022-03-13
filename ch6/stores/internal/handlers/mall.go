package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers application.DomainEventHandlers,
	domainSubscriber ddd.EventSubscriber,
) {
	domainSubscriber.Subscribe(domain.StoreCreated{}, mallHandlers.OnStoreCreated)
	domainSubscriber.Subscribe(domain.StoreParticipationEnabled{}, mallHandlers.OnStoreParticipationEnabled)
	domainSubscriber.Subscribe(domain.StoreParticipationDisabled{}, mallHandlers.OnStoreParticipationDisabled)
	domainSubscriber.Subscribe(domain.StoreRebranded{}, mallHandlers.OnStoreRebranded)
}
