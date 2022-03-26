package em

import (
	"context"
	"sync"

	"eda-in-golang/ch7/internal/ddd"
)

type Subscriber interface {
	ddd.EventSubscriber
	MessageHandler
}

type subscriber struct {
	handlers map[string]ddd.EventHandler
	mu       sync.Mutex
}

func NewSubscriber() Subscriber {
	return &subscriber{
		handlers: make(map[string]ddd.EventHandler),
	}
}

func (s *subscriber) Subscribe(name string, handler ddd.EventHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers[name] = handler
}

func (s *subscriber) HandleMessage(ctx context.Context, message Message) error {
	handler, exists := s.handlers[message.EventName()]
	if !exists {
		return nil
	}

	return handler.HandleEvent(ctx, message)
}
