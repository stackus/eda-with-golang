package es

import (
	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
)

type AggregateFactory func(id string) Aggregater

type Aggregater interface {
	AggregateName() string
}

type AggregateRoot struct {
	ID     string
	events []ddd.Event
}

func (a AggregateRoot) GetID() string {
	return a.ID
}

func (a *AggregateRoot) AddEvent(event ddd.Event) {
	a.events = append(a.events, event)
}

func (a AggregateRoot) GetEvents() []ddd.Event {
	return a.events
}
