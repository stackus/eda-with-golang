package domain

type Store struct {
	ID            string
	Name          string
	Location      Location
	Participating bool
}

func CreateStore(id, name string, location Location) (store *Store, err error) {
	if err = StoreNameCannotBeBlank(name); err != nil {
		return
	}

	if err = StoreLocationCannotBeBlank(location); err != nil {
		return
	}

	store = &Store{
		ID:       id,
		Name:     name,
		Location: location,
	}

	// TODO create StoreCreated integration event

	return
}

func (s *Store) EnableParticipation() (err error) {
	if err = StoreMustNotBeParticipating(s.Participating); err != nil {
		return err
	}

	s.Participating = true

	return nil
}
