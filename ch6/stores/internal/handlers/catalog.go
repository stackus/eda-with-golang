package handlers

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers application.DomainEventHandlers,
	domainSubscriber ddd.EventSubscriber,
) {
	domainSubscriber.Subscribe(domain.ProductAdded{}, catalogHandlers.OnProductAdded)
	domainSubscriber.Subscribe(domain.ProductRemoved{}, catalogHandlers.OnProductRemoved)
}
