package ddd

type AggregateNamer interface {
	AggregateName() string
}

type Eventer interface {
	AddEvent(string, EventPayload, ...EventOption)
	Events() []Event
	ClearEvents()
}

type Aggregate struct {
	Entity
	events []Event
}

var _ interface {
	AggregateNamer
	Eventer
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Entity: NewEntity(id, name),
		events: make([]Event, 0),
	}
}

func (a Aggregate) AggregateName() string { return a.name }
func (a Aggregate) Events() []Event       { return a.events }
func (a Aggregate) ClearEvents()          { a.events = []Event{} }

func (a *Aggregate) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(
		options,
		WithAggregateInfo(a.name, a.id),
	)
	a.events = append(
		a.events,
		NewEvent(name, payload, options...),
	)
}

func (a *Aggregate) setEvents(events []Event) { a.events = events }
