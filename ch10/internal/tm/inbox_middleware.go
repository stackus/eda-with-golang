package tm

import (
	"context"
	"fmt"

	"github.com/stackus/errors"

	"eda-in-golang/internal/am"
)

type ErrDuplicateMessage string

type InboxStore interface {
	Find(ctx context.Context, msgID string) (am.RawMessage, error)
	Save(ctx context.Context, msg am.RawMessage) error
}

type inbox struct {
	am.RawMessageStream
	store InboxStore
}

var _ am.RawMessageStream = (*inbox)(nil)

func NewInboxMiddleware(store InboxStore) am.RawMessageStreamMiddleware {
	o := inbox{store: store}

	return func(stream am.RawMessageStream) am.RawMessageStream {
		o.RawMessageStream = stream

		return o
	}
}

func (i inbox) Subscribe(topicName string, handler am.RawMessageHandler, options ...am.SubscriberOption) error {
	fn := am.MessageHandlerFunc[am.IncomingRawMessage](func(ctx context.Context, msg am.IncomingRawMessage) error {
		err := i.store.Save(ctx, msg)
		if err != nil {
			var errDupe ErrDuplicateMessage
			if errors.As(err, &errDupe) {
				// duplicate message; return without an error to let the message Ack
				return nil
			}
			return err
		}

		// try to insert the message
		return handler.HandleMessage(ctx, msg)

	})
	return i.RawMessageStream.Subscribe(topicName, fn, options...)
}

func (e ErrDuplicateMessage) Error() string {
	return fmt.Sprintf("duplicate message id encountered: %s", string(e))
}
