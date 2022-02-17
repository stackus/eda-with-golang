package domain

import (
	"fmt"
)

type CustomerID string

type Customer struct {
	ID        CustomerID
	Name      string
	SmsNumber string
	Enabled   bool
}

var (
	ErrNameCannotBeBlank       = fmt.Errorf("the customer name cannot be blank")
	ErrCustomerIDCannotBeBlank = fmt.Errorf("the customer id cannot be blank")
	ErrSmsNumberCannotBeBlank  = fmt.Errorf("the SMS number cannot be blank")
	ErrCustomerAlreadyEnabled  = fmt.Errorf("the customer is already enabled")
	ErrCustomerAlreadyDisabled = fmt.Errorf("the customer is already disabled")
)

func (i CustomerID) String() string {
	return string(i)
}

func ToCustomerID(id string) CustomerID {
	return CustomerID(id)
}

func RegisterCustomer(id CustomerID, name, smsNumber string) (*Customer, error) {
	if id == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if name == "" {
		return nil, ErrNameCannotBeBlank
	}

	if smsNumber == "" {
		return nil, ErrSmsNumberCannotBeBlank
	}

	return &Customer{
		ID:        id,
		Name:      name,
		SmsNumber: smsNumber,
		Enabled:   true,
	}, nil
}

func (c *Customer) Enable() error {
	if c.Enabled {
		return ErrCustomerAlreadyEnabled
	}

	c.Enabled = true

	return nil
}

func (c *Customer) Disable() error {
	if !c.Enabled {
		return ErrCustomerAlreadyDisabled
	}

	c.Enabled = false

	return nil
}
