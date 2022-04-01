package ddd

import (
	"time"
)

type EventOption func(*event)

func WithEventID(id string) EventOption {
	return func(e *event) {
		e.id = id
	}
}

func WithOccurredAt(occurredAt time.Time) EventOption {
	return func(e *event) {
		e.occurredAt = occurredAt
	}
}

func WithAggregateInfo(name, id string) EventOption {
	return func(e *event) {
		e.metadata.Set(AggregateIDKey, id)
		e.metadata.Set(AggregateNameKey, name)
	}
}

func WithAggregateVersion(version int) EventOption {
	return func(e *event) {
		e.metadata.Set(AggregateVersionKey, version)
	}
}
