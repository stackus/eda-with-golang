package logging

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/ordering/internal/application"
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

func (h DomainEventHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Ordering.%s.OnOrderCreated", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Ordering.%s.OnOrderCreated", h.label) }()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Ordering.%s.OnOrderReadied", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Ordering.%s.OnOrderReadied", h.label) }()
	return h.DomainEventHandlers.OnOrderReadied(ctx, event)
}

func (h DomainEventHandlers) OnOrderCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Ordering.%s.OnOrderCanceled", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Ordering.%s.OnOrderCanceled", h.label) }()
	return h.DomainEventHandlers.OnOrderCanceled(ctx, event)
}
