package ddd

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type (
	EventHandler func(ctx context.Context, event Event) error

	EventPayload interface {
		EventName() string
	}

	Event interface {
		ID() string
		Name() string
		Payload() EventPayload
		OccurredAt() time.Time
		AggregateName() string
		AggregateID() string
		AggregateVersion() int
		Metadata() map[string]interface{}
	}

	EventOption func(*event)

	event struct {
		id               string
		name             string
		payload          EventPayload
		occurredAt       time.Time
		aggregateName    string
		aggregateID      string
		aggregateVersion int
		metadata         map[string]interface{}
	}
)

var _ Event = (*event)(nil)

func NewEvent(payload EventPayload, options ...EventOption) Event {
	evt := event{
		id:         uuid.New().String(),
		name:       payload.EventName(),
		payload:    payload,
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option(&evt)
	}

	return evt
}

func (e event) ID() string {
	return e.id
}

func (e event) Name() string {
	return e.name
}

func (e event) Payload() EventPayload {
	return e.payload
}

func (e event) OccurredAt() time.Time {
	return e.occurredAt
}

func (e event) AggregateName() string {
	return e.aggregateName
}

func (e event) AggregateID() string {
	return e.aggregateID
}

func (e event) AggregateVersion() int {
	return e.aggregateVersion
}

func (e event) Metadata() map[string]interface{} {
	return e.metadata
}

func RegisterEventPayload(cd registry.Codec, payload EventPayload) error {
	return cd.Register(payload.EventName(), payload)
}
