package ddd

type Evented interface {
	AddEvent(EventPayload, ...EventOption)
	GetEvents() []Event
	ClearEvents()
}

type Aggregate struct {
	Entity
	events []Event
}

var _ interface {
	Evented
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Entity: NewEntity(id, name),
		events: make([]Event, 0),
	}
}

func (a Aggregate) ID() string {
	return a.id
}

func (a Aggregate) AggregateName() string {
	return a.name
}

func (a *Aggregate) AddEvent(payload EventPayload, options ...EventOption) {
	options = append(options, WithAggregateInfo(a.name, a.id))
	a.events = append(a.events, NewEvent(payload, options...))
}

func (a Aggregate) GetEvents() []Event {
	return a.events
}

func (a Aggregate) ClearEvents() {
	a.events = []Event{}
}
