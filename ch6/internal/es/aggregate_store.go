package es

import (
	"context"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type EventSourcedAggregate interface {
	ddd.IDed
	AggregateName() string
	ddd.Evented
	Versioned
	EventApplied
	EventCommitted
}

type AggregateStore interface {
	Load(ctx context.Context, aggregate EventSourcedAggregate) error
	Save(ctx context.Context, aggregate EventSourcedAggregate) error
}
