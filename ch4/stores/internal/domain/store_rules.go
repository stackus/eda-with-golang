package domain

import (
	"fmt"
)

var (
	ErrStoreNameIsBlank            = fmt.Errorf("a store name cannot be blank")
	ErrStoreLocationIsBlank        = fmt.Errorf("a store location cannot be blank")
	ErrStoreIsAlreadyParticipating = fmt.Errorf("the store is already participating")
)

func StoreNameCannotBeBlank(name string) error {
	if name == "" {
		return ErrStoreNameIsBlank
	}

	return nil
}

func StoreLocationCannotBeBlank(location Location) error {
	if location.String() == "" {
		return ErrStoreLocationIsBlank
	}

	return nil
}

func StoreMustNotBeParticipating(participating bool) error {
	if participating {
		return ErrStoreIsAlreadyParticipating
	}

	return nil
}
