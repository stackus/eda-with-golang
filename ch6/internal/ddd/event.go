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
		Occurred() time.Time
		AggregateName() string
		AggregateID() string
		Metadata() map[string]interface{}
	}

	EventOption func(*event)

	event struct {
		id            string
		name          string
		payload       EventPayload
		occurred      time.Time
		aggregateName string
		aggregateID   string
		metadata      map[string]interface{}
	}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, options ...EventOption) event {
	evt := event{
		id:       uuid.New().String(),
		name:     name,
		payload:  payload,
		occurred: time.Now(),
	}

	for _, option := range options {
		option(&evt)
	}

	return evt
}

func (e event) ID() string {
	return e.id
}

func (e event) EventName() string {
	return e.name
}

func (e event) Payload() EventPayload {
	return e.payload
}

func (e event) Occurred() time.Time {
	return e.occurred
}

func (e event) AggregateName() string {
	return e.aggregateName
}

func (e event) AggregateID() string {
	return e.aggregateID
}

func (e event) Metadata() map[string]interface{} {
	return e.metadata
}

func WithEventID(id string) EventOption {
	return func(e *event) {
		e.id = id
	}
}

func WithOccurred(occurred time.Time) EventOption {
	return func(e *event) {
		e.occurred = occurred
	}
}

func WithAggregateInfo(name, id string) EventOption {
	return func(e *event) {
		e.aggregateID = id
		e.aggregateName = name
	}
}
