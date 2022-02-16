package domain

type CustomerID string

type Customer struct {
	ID        CustomerID
	Name      string
	SmsNumber string
	Enabled   bool
}

func (i CustomerID) String() string {
	return string(i)
}

func RegisterCustomer(id CustomerID, name, smsNumber string) (*Customer, error) {
	// validate name

	// validate smsNumber

	return &Customer{
		ID:        id,
		Name:      name,
		SmsNumber: smsNumber,
		Enabled:   true,
	}, nil
}
