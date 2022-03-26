package handlers

import (
	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers ddd.EventHandler, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.StoreCreatedEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreParticipationEnabledEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreParticipationDisabledEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreRebrandedEvent, mallHandlers)
}
