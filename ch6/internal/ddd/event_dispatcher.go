package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(eventName string, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

var _ interface {
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (h *EventDispatcher) Subscribe(eventName string, handler EventHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[eventName] = append(h.handlers[eventName], handler)
}

func (h *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range h.handlers[event.EventName()] {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
