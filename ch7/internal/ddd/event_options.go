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

func WithOccurredAt(occurred time.Time) EventOption {
	return func(e *event) {
		e.occurredAt = occurred
	}
}

func WithAggregateInfo(name, id string) EventOption {
	return func(e *event) {
		e.aggID = id
		e.aggName = name
	}
}

func WithAggregateVersion(version int) EventOption {
	return func(e *event) {
		e.aggVersion = version
	}
}
