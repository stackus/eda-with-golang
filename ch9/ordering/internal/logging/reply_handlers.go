package logging

import (
	"context"

	"github.com/rs/zerolog"

	"eda-in-golang/ch9/internal/ac"
	"eda-in-golang/ch9/internal/ddd"
)

type sagaReplyHandlers[T any] struct {
	ac.Orchestrator[T]
	label  string
	logger zerolog.Logger
}

var _ ac.Orchestrator[any] = (*sagaReplyHandlers[any])(nil)

func LogReplyHandlerAccess[T any](orc ac.Orchestrator[T], label string, logger zerolog.Logger) ac.Orchestrator[T] {
	return sagaReplyHandlers[T]{
		Orchestrator: orc,
		label:        label,
		logger:       logger,
	}
}

func (h sagaReplyHandlers[T]) HandleReply(ctx context.Context, reply ddd.Reply) (err error) {
	h.logger.Info().Msgf("--> Ordering.%s.On(%s)", h.label, reply.ReplyName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- Ordering.%s.On(%s)", h.label, reply.ReplyName()) }()
	return h.Orchestrator.HandleReply(ctx, reply)
}
