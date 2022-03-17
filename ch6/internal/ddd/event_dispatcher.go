package ddd

import (
	"context"
	"sync"
)

type (
	EventHandler interface {
		HandleEvent(ctx context.Context, event Event) error
	}

	EventHandlerFunc func(ctx context.Context, event Event) error

	EventSubscriber interface {
		Subscribe(name string, handler EventHandler)
	}

	EventPublisher interface {
		Publish(ctx context.Context, events ...Event) error
	}

	EventDispatcher struct {
		handlers map[string][]EventHandler
		mu       sync.Mutex
	}
)

var _ interface {
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (h *EventDispatcher) Subscribe(name string, handler EventHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[name] = append(h.handlers[name], handler)
}

func (h *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range h.handlers[event.EventName()] {
			err := handler.HandleEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f EventHandlerFunc) HandleEvent(ctx context.Context, event Event) error {
	return f(ctx, event)
}
