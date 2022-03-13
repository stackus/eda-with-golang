package es

import (
	"context"
	"fmt"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type Versioned interface {
	Version() int
	PendingVersion() int
}

type EventApplied interface {
	ApplyEvent(context.Context, ddd.Event) error
}

type EventCommitted interface {
	CommitEvents()
}

type Aggregate struct {
	ddd.Aggregate
	version int
}

var _ interface {
	EventCommitted
	Versioned
	Versioner
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Aggregate: ddd.NewAggregate(id, name),
		version:   0,
	}
}

func (a *Aggregate) AddEvent(payload ddd.EventPayload, options ...ddd.EventOption) {
	options = append(options, ddd.WithAggregateVersion(a.PendingVersion()+1))
	a.Aggregate.AddEvent(payload, options...)
}

func (a *Aggregate) CommitEvents() {
	a.version += len(a.GetEvents())
	a.ClearEvents()
}

func (a Aggregate) Version() int {
	return a.version
}

func (a Aggregate) PendingVersion() int {
	return a.version + len(a.GetEvents())
}

func (a *Aggregate) setVersion(version int) {
	a.version = version
}

func LoadEvent(ctx context.Context, v interface{}, event ddd.Event) error {
	type loader interface {
		EventApplied
		Versioner
	}

	if agg, ok := v.(loader); !ok {
		return fmt.Errorf("%T does not have the method implemented to load events", v)
	} else {
		if err := agg.ApplyEvent(ctx, event); err != nil {
			return err
		}
		agg.setVersion(event.AggregateVersion())
	}
	return nil
}
