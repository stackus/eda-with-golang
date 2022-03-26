package em

import (
	"context"

	"eda-in-golang/ch7/internal/ddd"
)

type Headers map[string][]string

type MessageHandler interface {
	HandleMessage(ctx context.Context, message Message) error
}

type MessageHandlerFunc func(ctx context.Context, message Message) error

type MessageSubscriber interface {
	Subscribe(name string, handler MessageHandler)
}

type MessagePublisher interface {
	Publish(ctx context.Context, message Message) error
}

type Message interface {
	ddd.Event
	Headers() Headers
	Ack() error
	NAck() error
	Extend() error
	Kill() error
}

type message struct {
	ddd.Event
	headers  Headers
	acked    bool
	ackFn    func() error
	nackFn   func() error
	extendFn func() error
	killFn   func() error
}

func NewMessage(event ddd.Event, options ...MessageOption) Message {
	m := &message{
		Event:   event,
		headers: make(Headers),
	}

	for _, option := range options {
		option.configureMessage(m)
	}

	return m
}

func (m *message) Headers() Headers {
	return m.headers
}

func (m *message) Ack() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.ackFn()
}

func (m *message) NAck() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.nackFn()
}

func (m message) Extend() error {
	return m.extendFn()
}

func (m *message) Kill() error {
	if m.acked {
		return nil
	}

	m.acked = true
	return m.killFn()
}

func (h Headers) Add(key, value string) {
	h[key] = append(h[key], value)
}

func (h Headers) Set(key, value string) {
	h[key] = []string{value}
}

func (h Headers) Get(key string) string {
	if h == nil {
		return ""
	}
	if v := h[key]; v != nil {
		return v[0]
	}
	return ""
}

func (h Headers) Values(key string) []string {
	return h[key]
}

func (h Headers) Del(key string) {
	delete(h, key)
}
