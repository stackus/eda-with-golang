package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	EventPayload interface{}

	Event interface {
		IDer
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	event struct {
		Entity
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Event = (*event)(nil)

func newEvent(name string, payload EventPayload, options ...EventOption) event {
	evt := event{
		Entity:     NewEntity(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option(&evt)
	}

	return evt
}

func (e event) EventName() string     { return e.name }
func (e event) Payload() EventPayload { return e.payload }
func (e event) Metadata() Metadata    { return e.metadata }
func (e event) OccurredAt() time.Time { return e.occurredAt }
