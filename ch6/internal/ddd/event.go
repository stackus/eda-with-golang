package ddd

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	EventHandler func(ctx context.Context, event Event) error

	EventPayload interface{}

	Event interface {
		ID() string
		EventName() string
		Payload() EventPayload
		OccurredAt() time.Time
		AggregateName() string
		AggregateID() string
		AggregateVersion() int
	}

	EventOption func(*event)

	event struct {
		id         string
		name       string
		payload    EventPayload
		occurredAt time.Time
		aggName    string
		aggID      string
		aggVersion int
	}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, options ...EventOption) Event {
	evt := event{
		id:         uuid.New().String(),
		name:       name,
		payload:    payload,
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option(&evt)
	}

	return evt
}

func (e event) ID() string            { return e.id }
func (e event) EventName() string     { return e.name }
func (e event) Payload() EventPayload { return e.payload }
func (e event) OccurredAt() time.Time { return e.occurredAt }
func (e event) AggregateName() string { return e.aggName }
func (e event) AggregateID() string   { return e.aggID }
func (e event) AggregateVersion() int { return e.aggVersion }
