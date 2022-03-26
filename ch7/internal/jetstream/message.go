package jetstream

import (
	"time"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/em"
	"eda-in-golang/ch7/internal/registry"
)

type message struct {
	reg        registry.Registry
	eventID    string
	eventName  string
	data       []byte
	payload    ddd.EventPayload
	aggID      string
	aggName    string
	aggVersion int
	occurredAt time.Time
	headers    em.Headers
	acked      bool
	ackFn      func() error
	nackFn     func() error
	extendFn   func() error
	killFn     func() error
}

var _ em.Message = (*message)(nil)

func (m message) ID() string {
	return m.eventID
}

func (m message) EventName() string {
	return m.eventName
}

func (m *message) Payload() ddd.EventPayload {
	var err error

	if m.payload != nil {
		return m.payload
	}

	m.payload, err = m.reg.Deserialize(m.eventName, m.data)
	if err != nil {
		m.payload = err
	}

	return m.payload
}

func (m message) OccurredAt() time.Time {
	return m.occurredAt
}

func (m message) AggregateName() string {
	return m.aggName
}

func (m message) AggregateID() string {
	return m.aggID
}

func (m message) AggregateVersion() int {
	return m.aggVersion
}

func (m message) Headers() em.Headers {
	return m.headers
}

func (m *message) Ack() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.ackFn()
}

func (m *message) NAck() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.nackFn()
}

func (m message) Extend() error {
	return m.extendFn()
}

func (m *message) Kill() error {
	if m.acked {
		return nil
	}

	m.acked = true
	return m.killFn()
}
