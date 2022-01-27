package domain

import (
	"fmt"
)

var (
	ErrOfferingNameIsBlank     = fmt.Errorf("the offering name cannot be blank")
	ErrOfferingPriceIsNegative = fmt.Errorf("the offering price cannot be negative")
)

type Offering struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	Price       float64
}

func CreateOffering(id, storeID, name, description string, price float64) (*Offering, error) {
	if name == "" {
		return nil, ErrOfferingNameIsBlank
	}

	if price < 0 {
		return nil, ErrOfferingPriceIsNegative
	}

	offering := &Offering{
		ID:          id,
		StoreID:     storeID,
		Name:        name,
		Description: description,
		Price:       price,
	}

	return offering, nil
}
