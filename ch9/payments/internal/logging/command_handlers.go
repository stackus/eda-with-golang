package logging

import (
	"context"

	"github.com/rs/zerolog"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
)

type CommandHandlers[T ddd.Command] struct {
	am.CommandHandler
	label  string
	logger zerolog.Logger
}

var _ am.CommandHandler = (*CommandHandlers[ddd.Command])(nil)

func LogCommandHandlerAccess[T ddd.Command](handlers am.CommandHandler, label string, logger zerolog.Logger) am.CommandHandler {
	return CommandHandlers[T]{
		CommandHandler: handlers,
		label:          label,
		logger:         logger,
	}
}

func (h CommandHandlers[T]) HandleCommand(ctx context.Context, command ddd.Command) (reply ddd.Reply, err error) {
	h.logger.Info().Msgf("--> Payments.%s.On(%s)", h.label, command.CommandName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- Payments.%s.On(%s)", h.label, command.CommandName()) }()
	return h.CommandHandler.HandleCommand(ctx, command)
}