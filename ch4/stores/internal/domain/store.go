package domain

import (
	"fmt"
)

var (
	ErrStoreNameIsBlank               = fmt.Errorf("the store name cannot be blank")
	ErrStoreLocationIsBlank           = fmt.Errorf("the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = fmt.Errorf("the store is already participating")
	ErrStoreIsAlreadyNotParticipating = fmt.Errorf("the store is already not participating")
)

type Store struct {
	ID            string
	Name          string
	Location      Location
	Participating bool
}

func CreateStore(id, name string, location Location) (store *Store, err error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store = &Store{
		ID:       id,
		Name:     name,
		Location: location,
	}

	return
}

func (s *Store) EnableParticipation() (err error) {
	if s.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	s.Participating = true

	return
}

func (s *Store) DisableParticipation() (err error) {
	if !s.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	s.Participating = false

	return
}
