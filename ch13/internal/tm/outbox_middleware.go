package tm

import (
	"context"

	"github.com/stackus/errors"

	"eda-in-golang/internal/am"
)

type OutboxStore interface {
	Save(ctx context.Context, msg am.Message) error
	FindUnpublished(ctx context.Context, limit int) ([]am.Message, error)
	MarkPublished(ctx context.Context, ids ...string) error
}

type outbox struct {
	am.MessageStream
	store OutboxStore
}

var _ am.MessageStream = (*outbox)(nil)

func NewOutboxStreamMiddleware(store OutboxStore) am.MessageStreamMiddleware {
	o := outbox{store: store}

	return func(stream am.MessageStream) am.MessageStream {
		o.MessageStream = stream

		return o
	}
}

func (o outbox) Publish(ctx context.Context, topicName string, msg am.Message) error {
	err := o.store.Save(ctx, msg)

	var errDupe ErrDuplicateMessage
	if errors.As(err, &errDupe) {
		return nil
	}

	return err
}

func OutboxPublisher(store OutboxStore) am.MessagePublisherMiddleware {
	return func(next am.MessagePublisher) am.MessagePublisher {
		return am.MessagePublisherFunc(func(ctx context.Context, topicName string, msg am.Message) error {
			err := store.Save(ctx, msg)
			var errDupe ErrDuplicateMessage
			if errors.As(err, &errDupe) {
				return nil
			}
			return err
		})
	}
}
