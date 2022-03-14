package logging

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/stores/internal/application"
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

func (h DomainEventHandlers) OnStoreCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnStoreCreated", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnStoreCreated", h.label) }()
	return h.DomainEventHandlers.OnStoreCreated(ctx, event)
}

func (h DomainEventHandlers) OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnStoreParticipationEnabled", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnStoreParticipationEnabled", h.label) }()
	return h.DomainEventHandlers.OnStoreParticipationEnabled(ctx, event)
}

func (h DomainEventHandlers) OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnStoreParticipationDisabled", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnStoreParticipationDisabled", h.label) }()
	return h.DomainEventHandlers.OnStoreParticipationDisabled(ctx, event)
}

func (h DomainEventHandlers) OnStoreRebranded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnStoreRebranded", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnStoreRebranded", h.label) }()
	return h.DomainEventHandlers.OnStoreRebranded(ctx, event)
}

func (h DomainEventHandlers) OnProductAdded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnProductAdded", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnProductAdded", h.label) }()
	return h.DomainEventHandlers.OnProductAdded(ctx, event)
}

func (h DomainEventHandlers) OnProductRebranded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnProductRebranded", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnProductRebranded", h.label) }()
	return h.DomainEventHandlers.OnProductRebranded(ctx, event)
}

func (h DomainEventHandlers) OnProductPriceIncreased(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnProductPriceIncreased", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnProductPriceIncreased", h.label) }()
	return h.DomainEventHandlers.OnProductPriceIncreased(ctx, event)
}

func (h DomainEventHandlers) OnProductPriceDecreased(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnProductPriceDecreased", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnProductPriceDecreased", h.label) }()
	return h.DomainEventHandlers.OnProductPriceDecreased(ctx, event)
}

func (h DomainEventHandlers) OnProductRemoved(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.OnProductRemoved", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.OnProductRemoved", h.label) }()
	return h.DomainEventHandlers.OnProductRemoved(ctx, event)
}
