package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers application.DomainEventHandlers,
	domainSubscriber ddd.EventSubscriber,
) {
	domainSubscriber.Subscribe(domain.StoreCreatedEvent, mallHandlers.OnStoreCreated)
	domainSubscriber.Subscribe(domain.StoreParticipationEnabledEvent, mallHandlers.OnStoreParticipationEnabled)
	domainSubscriber.Subscribe(domain.StoreParticipationDisabledEvent, mallHandlers.OnStoreParticipationDisabled)
	domainSubscriber.Subscribe(domain.StoreRebrandedEvent, mallHandlers.OnStoreRebranded)
}
