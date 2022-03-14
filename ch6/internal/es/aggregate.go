package es

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type Versioner interface {
	Version() int
	PendingVersion() int
}

type Aggregate struct {
	ddd.Aggregate
	version int
}

var _ interface {
	EventCommitter
	Versioner
	VersionSetter
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Aggregate: ddd.NewAggregate(id, name),
		version:   0,
	}
}

func (a *Aggregate) AddEvent(name string, payload ddd.EventPayload, options ...ddd.EventOption) {
	options = append(options, ddd.WithAggregateVersion(a.PendingVersion()+1))
	a.Aggregate.AddEvent(name, payload, options...)
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
