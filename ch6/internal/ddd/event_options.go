package ddd

import (
	"time"
)

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
		e.aggregateID = id
		e.aggregateName = name
	}
}

func WithAggregateVersion(version int) EventOption {
	return func(e *event) {
		e.aggregateVersion = version
	}
}
