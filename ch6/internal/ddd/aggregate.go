package ddd

type Aggregate struct {
	id      string
	aggName string
	events  []Event
}

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		id:      id,
		aggName: name,
		events:  make([]Event, 0),
	}
}

func (a Aggregate) ID() string {
	return a.id
}

func (a *Aggregate) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(options, WithAggregateInfo(a.aggName, a.id))
	a.events = append(a.events, NewEvent(name, payload, options...))
}

func (a Aggregate) GetEvents() []Event {
	return a.events
}

func (a Aggregate) ClearEvents() {
	a.events = []Event{}
}

func (a *Aggregate) setID(id string) {
	a.id = id
}

func (a *Aggregate) setAggregateName(name string) {
	a.aggName = name
}
