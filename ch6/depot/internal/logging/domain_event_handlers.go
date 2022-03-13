package logging

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/stackus/eda-with-golang/ch6/depot/internal/application"
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	label  string
	logger zerolog.Logger
}

var _ application.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers application.DomainEventHandlers, label string, logger zerolog.Logger,
) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		label:               label,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnShoppingListCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Depot.%s.OnShoppingListCreated", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Depot.%s.OnShoppingListCreated", h.label) }()
	return h.DomainEventHandlers.OnShoppingListCreated(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Depot.%s.OnShoppingListCanceled", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Depot.%s.OnShoppingListCanceled", h.label) }()
	return h.DomainEventHandlers.OnShoppingListCanceled(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListAssigned(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Depot.%s.OnShoppingListAssigned", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Depot.%s.OnShoppingListAssigned", h.label) }()
	return h.DomainEventHandlers.OnShoppingListAssigned(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Depot.%s.OnShoppingListCompleted", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Depot.%s.OnShoppingListCompleted", h.label) }()
	return h.DomainEventHandlers.OnShoppingListCompleted(ctx, event)
}
