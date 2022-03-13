package domain

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/es"
)

const StoreAggregate = "stores.Store"

var (
	ErrStoreNameIsBlank               = errors.Wrap(errors.ErrBadRequest, "the store name cannot be blank")
	ErrStoreLocationIsBlank           = errors.Wrap(errors.ErrBadRequest, "the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "the store is already participating")
	ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "the store is already not participating")
)

type Store struct {
	es.Aggregate
	Name          string
	Location      string
	Participating bool
}

func NewStore(id string) *Store {
	return &Store{
		Aggregate: es.NewAggregate(id, StoreAggregate),
	}
}

func CreateStore(id, name, location string) (*Store, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store := NewStore(id)
	store.Name = name
	store.Location = location

	store.AddEvent(&StoreCreated{
		Name:     name,
		Location: location,
	})

	return store, nil
}

func (s *Store) EnableParticipation() (err error) {
	if s.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	s.AddEvent(&StoreParticipationEnabled{})

	return
}

func (s *Store) DisableParticipation() (err error) {
	if !s.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	s.AddEvent(&StoreParticipationDisabled{})

	return
}

func (s *Store) Rebrand(name string) error {
	s.AddEvent(&StoreRebranded{
		Name: name,
	})

	return nil
}

func (s *Store) ApplyEvent(_ context.Context, event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *StoreCreated:
		s.Name = payload.Name
		s.Location = payload.Location

	case *StoreParticipationEnabled:
		s.Participating = true

	case *StoreParticipationDisabled:
		s.Participating = false

	case *StoreRebranded:
		s.Name = payload.Name

	default:
		return errors.ErrInternal.Msgf("%T received the expected event payload %T", s, event.Payload())
	}

	return nil
}
