package logging

import (
	"context"

	"github.com/rs/zerolog"

	"eda-in-golang/ch7/internal/ddd"
)

type EventHandlers struct {
	ddd.EventHandler
	label  string
	logger zerolog.Logger
}

var _ ddd.EventHandler = (*EventHandlers)(nil)

func LogEventHandlerAccess(handlers ddd.EventHandler, label string, logger zerolog.Logger) EventHandlers {
	return EventHandlers{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers) HandleEvent(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
