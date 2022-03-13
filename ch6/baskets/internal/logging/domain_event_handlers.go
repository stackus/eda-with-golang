package logging

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/stackus/eda-with-golang/ch6/baskets/internal/application"
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

func (h DomainEventHandlers) OnBasketStarted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.OnBasketStarted", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.OnBasketStarted", h.label) }()
	return h.DomainEventHandlers.OnBasketStarted(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemAdded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.OnBasketItemAdded", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.OnBasketItemAdded", h.label) }()
	return h.DomainEventHandlers.OnBasketItemAdded(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemRemoved(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.OnBasketItemRemoved", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.OnBasketItemRemoved", h.label) }()
	return h.DomainEventHandlers.OnBasketItemRemoved(ctx, event)
}

func (h DomainEventHandlers) OnBasketCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.OnBasketCanceled", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.OnBasketCanceled", h.label) }()
	return h.DomainEventHandlers.OnBasketCanceled(ctx, event)
}

func (h DomainEventHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.OnBasketCheckedOut", h.label)
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.OnBasketCheckedOut", h.label) }()
	return h.DomainEventHandlers.OnBasketCheckedOut(ctx, event)
}
